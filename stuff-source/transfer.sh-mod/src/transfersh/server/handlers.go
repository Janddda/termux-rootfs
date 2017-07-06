/*
The MIT License (MIT)

Copyright (c) 2014-2017 DutchCoders [https://github.com/dutchcoders/]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package server

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	html_template "html/template"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"mime"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	text_template "text/template"
	"time"

	web "transfersh/frontend"
	"transfersh/server/codec"
	"transfersh/server/render"

	"github.com/golang/gddo/httputil/header"
	"github.com/gorilla/mux"
)

var (
	htmlTemplates = initHTMLTemplates()
	textTemplates = initTextTemplates()
)

func initTextTemplates() *text_template.Template {
	templateMap := text_template.FuncMap{"format": render.RenderFloat}

	// Templates with functions available to them
	var templates = text_template.New("").Funcs(templateMap)
	return templates
}

func initHTMLTemplates() *html_template.Template {
	templateMap := html_template.FuncMap{"format": render.RenderFloat}

	// Templates with functions available to them
	var templates = html_template.New("").Funcs(templateMap)

	return templates
}

func stripPrefix(path string) string {
	return strings.Replace(path, web.Prefix+"/", "", -1)
}

func sanitize(fileName string) string {
	return path.Clean(path.Base(fileName))
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

// Request.RemoteAddress contains port, which we want to remove i.e.:
// "[::1]:58292" => "[::1]"
func ipAddrFromRemoteAddr(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}

func acceptsHTML(hdr http.Header) bool {
	actual := header.ParseAccept(hdr, "Accept")

	for _, s := range actual {
		if s.Value == "text/html" {
			return (true)
		}
	}

	return (false)
}

func getURL(r *http.Request) *url.URL {
	u := *r.URL

	if r.TLS != nil {
		u.Scheme = "https"
	} else if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		u.Scheme = proto
	} else {
		u.Scheme = "http"
	}

	if u.Host != "" {
	} else if host, port, err := net.SplitHostPort(r.Host); err != nil {
		u.Host = r.Host
	} else {
		if port == "80" && u.Scheme == "http" {
			u.Host = host
		} else if port == "443" && u.Scheme == "https" {
			u.Host = host
		} else {
			u.Host = net.JoinHostPort(host, port)
		}
	}

	return &u
}

func (s *Server) notFoundError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	if acceptsHTML(r.Header) {
		if err := htmlTemplates.ExecuteTemplate(w, "404.html", nil); err != nil {
			log.Printf("notFoundError() - %s", err.Error())
			w.Write([]byte("Requested content is not found. Sorry...\n"))
			return
		}
	} else {
		w.Write([]byte("Requested content is not found. Sorry...\n"))
	}
}

func (s *Server) internalServerError(w http.ResponseWriter, r *http.Request, info string) {
	log.Println(info)

	w.WriteHeader(http.StatusInternalServerError)

	if acceptsHTML(r.Header) {
		if err := htmlTemplates.ExecuteTemplate(w, "500.html", nil); err != nil {
			log.Printf("internalServerError() - %s", err.Error())
			w.Write([]byte("Oops! Something went wrong...\n"))
			return
		}
	} else {
		w.Write([]byte("Oops! Something went wrong...\n"))
	}
}

// Redirect from http to https
func (s *Server) redirectHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.forceHTTPs {
			// we don't want to enforce https
		} else if strings.HasSuffix(ipAddrFromRemoteAddr(r.Host), ".onion") {
			// .onion addresses cannot get a valid certificate, so don't redirect
		} else if getURL(r).Scheme == "https" {
		} else {
			u := getURL(r)
			u.Scheme = "https"

			http.Redirect(w, r, u.String(), http.StatusPermanentRedirect)
			return
		}

		h.ServeHTTP(w, r)
	}
}

// Main (index) page
func (s *Server) mainPage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		SiteURL string
	}{
		getURL(r).String(),
	}

	if acceptsHTML(r.Header) {
		if err := htmlTemplates.ExecuteTemplate(w, "index.html", data); err != nil {
			s.internalServerError(w, r, fmt.Sprintf("mainPage() - %s", err.Error()))
			return
		}
	} else {
		if err := textTemplates.ExecuteTemplate(w, "index.txt", data); err != nil {
			s.internalServerError(w, r, fmt.Sprintf("mainPage() - %s", err.Error()))
			return
		}
	}
}

// File preview (download) page
func (s *Server) previewPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	filename := vars["filename"]

	contentType, contentLength, err := s.storage.Head(token, filename)
	if err != nil {
		s.notFoundError(w, r)
		return
	}

	// Convert bytes to KiB, MiB or GiB
	lengthUnit := "B"
	var fileLength float64

	if contentLength >= 1024 && contentLength < 1048576 {
		lengthUnit = "KiB"
		fileLength = toFixed(float64(contentLength)/1024, 3)
	} else if contentLength >= 1048576 && contentLength < 1073741824 {
		lengthUnit = "MiB"
		fileLength = toFixed(float64(contentLength)/1048576, 3)
	} else if contentLength >= 1073741824 {
		lengthUnit = "GiB"
		fileLength = toFixed(float64(contentLength)/1073741824, 3)
	} else {
		fileLength = float64(contentLength)
	}

	data := struct {
		ContentType string
		Filename    string
		Url         string
		FileLength  float64
		LengthUnit  string
	}{
		contentType,
		filename,
		r.URL.String(),
		fileLength,
		lengthUnit,
	}

	if err := htmlTemplates.ExecuteTemplate(w, "download.html", data); err != nil {
		s.internalServerError(w, r, fmt.Sprintf("previewPage() - %s", err.Error()))
		return
	}

}

// Process file uploaded via POST request
func (s *Server) postFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(_24K); nil != err {
		s.internalServerError(w, r, fmt.Sprintf("postFile() - %s", err.Error()))
		return
	}

	token := codec.Encode(10000000 + int64(rand.Intn(1000000000)))

	w.Header().Set("Content-Type", "text/plain")

	for _, fheaders := range r.MultipartForm.File {
		for _, fheader := range fheaders {
			filename := sanitize(fheader.Filename)
			contentType := fheader.Header.Get("Content-Type")

			if contentType == "" {
				contentType = mime.TypeByExtension(filepath.Ext(fheader.Filename))
			}

			var f io.Reader
			var err error

			if f, err = fheader.Open(); err != nil {
				s.internalServerError(w, r, fmt.Sprintf("postFile() - %s", err.Error()))
				return
			}

			var b bytes.Buffer

			n, err := io.CopyN(&b, f, _24K+1)
			if err != nil && err != io.EOF {
				s.internalServerError(w, r, fmt.Sprintf("postFile() - %s", err.Error()))
				return
			}

			var reader io.Reader

			if n > _24K {
				file, err := ioutil.TempFile(s.tempPath, "transfer-")
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()

				n, err = io.Copy(file, io.MultiReader(&b, f))
				if err != nil {
					os.Remove(file.Name())
					s.internalServerError(w, r, fmt.Sprintf("postFile() - %s", err.Error()))
					return
				}

				reader, err = os.Open(file.Name())
			} else {
				reader = bytes.NewReader(b.Bytes())
			}

			contentLength := n

			metadata := metadataForRequest(contentType, r)

			buffer := &bytes.Buffer{}
			if err := json.NewEncoder(buffer).Encode(metadata); err != nil {
				s.internalServerError(w, r, fmt.Sprintf("postFile() - %s", err.Error()))
				return
			} else if err := s.storage.Put(token, fmt.Sprintf("%s.metadata", filename), buffer, "text/json", uint64(buffer.Len())); err != nil {
				s.internalServerError(w, r, fmt.Sprintf("postFile() - %s", err.Error()))
				return
			}

			log.Printf("uploading %s \"%s\" %d \"%s\"", token, filename, contentLength, contentType)

			if err = s.storage.Put(token, filename, reader, contentType, uint64(contentLength)); err != nil {
				s.internalServerError(w, r, fmt.Sprintf("postFile() - %s", err.Error()))
				return

			}

			relativeURL, _ := url.Parse(path.Join(token, filename))
			fmt.Fprint(w, getURL(r).ResolveReference(relativeURL).String(), "\n")
		}
	}
}

// Process file uploaded via PUT request
func (s *Server) putFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	filename := sanitize(vars["filename"])

	contentLength := r.ContentLength

	var reader io.Reader

	reader = r.Body

	if contentLength == -1 {
		// queue file to disk, because s3 needs content length
		var err error
		var f io.Reader

		f = reader

		var b bytes.Buffer

		n, err := io.CopyN(&b, f, _24K+1)
		if err != nil && err != io.EOF {
			s.internalServerError(w, r, fmt.Sprintf("putFile() - %s", err.Error()))
			return
		}

		if n > _24K {
			file, err := ioutil.TempFile(s.tempPath, "transfer-")
			if err != nil {
				s.internalServerError(w, r, fmt.Sprintf("putFile() - %s", err.Error()))
				return
			}

			defer file.Close()

			n, err = io.Copy(file, io.MultiReader(&b, f))
			if err != nil {
				os.Remove(file.Name())
				s.internalServerError(w, r, fmt.Sprintf("putFile() - %s", err.Error()))
				return
			}

			reader, err = os.Open(file.Name())
		} else {
			reader = bytes.NewReader(b.Bytes())
		}

		contentLength = n
	}

	contentType := r.Header.Get("Content-Type")

	if contentType == "" {
		contentType = mime.TypeByExtension(filepath.Ext(vars["filename"]))
	}

	token := codec.Encode(10000000 + int64(rand.Intn(1000000000)))

	metadata := metadataForRequest(contentType, r)

	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(metadata); err != nil {
		s.internalServerError(w, r, fmt.Sprintf("putFile() - %s", err.Error()))
		return
	} else if err := s.storage.Put(token, fmt.Sprintf("%s.metadata", filename), buffer, "text/json", uint64(buffer.Len())); err != nil {
		s.internalServerError(w, r, fmt.Sprintf("putFile() - %s", err.Error()))
		return
	}

	log.Printf("uploading %s \"%s\" %d \"%s\"", token, filename, contentLength, contentType)

	var err error

	if err = s.storage.Put(token, filename, reader, contentType, uint64(contentLength)); err != nil {
		s.internalServerError(w, r, fmt.Sprintf("putFile() - %s", err.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	relativeURL, _ := url.Parse(path.Join(token, filename))
	fmt.Fprint(w, getURL(r).ResolveReference(relativeURL).String(), "\n")
}

// Send file to client
func (s *Server) sendFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	token := vars["token"]
	filename := vars["filename"]

	reader, contentType, contentLength, err := s.storage.Get(token, filename)
	if s.storage.IsNotExist(err) {
		s.notFoundError(w, r)
		return
	} else if err != nil {
		s.internalServerError(w, r, fmt.Sprintf("sendFile() - %s", err.Error()))
		return
	}

	if err := s.checkAgeExpiry(token, filename); err != nil {
		log.Println(err.Error())
		s.notFoundError(w, r)
		return
	}

	defer reader.Close()

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", strconv.FormatUint(contentLength, 10))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	httpVersion, _, _ := http.ParseHTTPVersion(r.Proto)
	if httpVersion < 2 {
		w.Header().Set("Connection", "close")
	}

	if _, err = io.Copy(w, reader); err != nil {
		s.internalServerError(w, r, fmt.Sprintf("sendFile() - %s", err.Error()))
		return
	}

	if err := s.checkDownloadsExpiry(token, filename); err != nil {
		log.Println(err.Error())
	}
}

// Archive files with gzip compressed tar and send to client
func (s *Server) sendTarGz(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	files := vars["files"]

	tarfilename := fmt.Sprintf("transfersh-%d.tar.gz", uint16(time.Now().UnixNano()))

	httpVersion, _, _ := http.ParseHTTPVersion(r.Proto)
	if httpVersion < 2 {
		w.Header().Set("Connection", "close")
	}

	// Check files in separate loop before processing
	for _, key := range strings.Split(files, ",") {
		if strings.HasPrefix(key, "/") {
			key = key[1:]
		}

		key = strings.Replace(key, "\\", "/", -1)

		token := strings.Split(key, "/")[0]
		filename := sanitize(strings.Split(key, "/")[1])

		if err := s.checkAgeExpiry(token, filename); err != nil {
			log.Println(err.Error())
		}

		reader, _, _, err := s.storage.Get(token, filename)
		if err != nil {
			if s.storage.IsNotExist(err) {
				s.notFoundError(w, r)
				return
			} else {
				s.internalServerError(w, r, fmt.Sprintf("sendTarGz() - %s", err.Error()))
				return
			}
		}
		reader.Close()
	}

	w.Header().Set("Content-Type", "application/x-gzip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", tarfilename))

	os := gzip.NewWriter(w)
	defer os.Close()

	zw := tar.NewWriter(os)
	defer zw.Close()

	for _, key := range strings.Split(files, ",") {
		if strings.HasPrefix(key, "/") {
			key = key[1:]
		}

		key = strings.Replace(key, "\\", "/", -1)

		token := strings.Split(key, "/")[0]
		filename := sanitize(strings.Split(key, "/")[1])

		if err := s.checkAgeExpiry(token, filename); err != nil {
			log.Println(err.Error())
			continue
		}

		reader, _, contentLength, err := s.storage.Get(token, filename)
		if err != nil {
			if s.storage.IsNotExist(err) {
				s.notFoundError(w, r)
				return
			} else {
				s.internalServerError(w, r, fmt.Sprintf("sendTarGz() - %s", err.Error()))
				return
			}
		}

		defer reader.Close()

		header := &tar.Header{
			Name: strings.Split(key, "/")[1],
			Size: int64(contentLength),
		}

		err = zw.WriteHeader(header)
		if err != nil {
			s.internalServerError(w, r, fmt.Sprintf("sendTarGz() - %s", err.Error()))
			return
		}

		if _, err = io.Copy(zw, reader); err != nil {
			s.internalServerError(w, r, fmt.Sprintf("sendTarGz() - %s", err.Error()))
			return
		}

		if err := s.checkDownloadsExpiry(token, filename); err != nil {
			log.Println(err.Error())
		}
	}
}

// Archive files with zip and send to client
func (s *Server) sendZip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	files := vars["files"]

	zipfilename := fmt.Sprintf("transfersh-%d.zip", uint16(time.Now().UnixNano()))

	httpVersion, _, _ := http.ParseHTTPVersion(r.Proto)
	if httpVersion < 2 {
		w.Header().Set("Connection", "close")
	}

	// Check files in separate loop before processing
	for _, key := range strings.Split(files, ",") {
		if strings.HasPrefix(key, "/") {
			key = key[1:]
		}

		key = strings.Replace(key, "\\", "/", -1)

		token := strings.Split(key, "/")[0]
		filename := sanitize(strings.Split(key, "/")[1])

		if err := s.checkAgeExpiry(token, filename); err != nil {
			log.Println(err.Error())
		}

		reader, _, _, err := s.storage.Get(token, filename)
		if err != nil {
			if s.storage.IsNotExist(err) {
				s.notFoundError(w, r)
				return
			} else {
				s.internalServerError(w, r, fmt.Sprintf("sendZip() - %s", err.Error()))
				return
			}
		}
		reader.Close()
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipfilename))

	zw := zip.NewWriter(w)

	for _, key := range strings.Split(files, ",") {
		if strings.HasPrefix(key, "/") {
			key = key[1:]
		}

		key = strings.Replace(key, "\\", "/", -1)

		token := strings.Split(key, "/")[0]
		filename := sanitize(strings.Split(key, "/")[1])

		reader, _, _, err := s.storage.Get(token, filename)
		if err != nil {
			if s.storage.IsNotExist(err) {
				s.notFoundError(w, r)
				return
			} else {
				s.internalServerError(w, r, fmt.Sprintf("sendZip() - %s", err.Error()))
				return
			}
		}

		defer reader.Close()

		header := &zip.FileHeader{
			Name:         strings.Split(key, "/")[1],
			Method:       zip.Store,
			ModifiedTime: uint16(time.Now().UnixNano()),
			ModifiedDate: uint16(time.Now().UnixNano()),
		}

		fw, err := zw.CreateHeader(header)

		if err != nil {
			s.internalServerError(w, r, fmt.Sprintf("sendZip() - %s", err.Error()))
			return
		}

		if _, err = io.Copy(fw, reader); err != nil {
			s.internalServerError(w, r, fmt.Sprintf("sendZip() - %s", err.Error()))
			return
		}

		if err := s.checkDownloadsExpiry(token, filename); err != nil {
			log.Println(err.Error())
		}
	}

	if err := zw.Close(); err != nil {
		s.internalServerError(w, r, fmt.Sprintf("sendZip() - %s", err.Error()))
		return
	}
}
