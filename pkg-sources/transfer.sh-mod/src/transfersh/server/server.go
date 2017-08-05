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
	_ "net/http/pprof"

	"crypto/tls"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/PuerkitoBio/ghost/handlers"
	"github.com/VojtechVitek/ratelimit"
	"github.com/VojtechVitek/ratelimit/memory"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/mux"

	web "transfersh/frontend"
	"transfersh/server/mime_types"
)

const SERVER_INFO = "transfer.sh"

// parse request with maximum memory of _24Kilobits
const _24K = (1 << 20) * 24

type OptionFn func(*Server)

type Server struct {
	tlsConfig *tls.Config

	profilerEnabled bool

	locks map[string]*sync.Mutex

	rateLimitRequests int

	storage Storage

	forceHTTPs bool

	tempPath string

	ListenerString    string
	TLSListenerString string

	Certificate string
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	mime_types.Setup()
}

func New(options ...OptionFn) (*Server, error) {
	s := &Server{
		locks: map[string]*sync.Mutex{},
	}

	for _, optionFn := range options {
		optionFn(s)
	}

	return s, nil
}

func Listener(s string) OptionFn {
	return func(srvr *Server) {
		srvr.ListenerString = s
	}
}

func TLSListener(s string) OptionFn {
	return func(srvr *Server) {
		srvr.TLSListenerString = s
	}
}

func ForceHTTPs() OptionFn {
	return func(srvr *Server) {
		srvr.forceHTTPs = true
	}
}

func TempPath(s string) OptionFn {
	return func(srvr *Server) {
		srvr.tempPath = s
	}
}

func RateLimit(requests int) OptionFn {
	return func(srvr *Server) {
		srvr.rateLimitRequests = requests
	}
}

func EnableProfiler() OptionFn {
	return func(srvr *Server) {
		srvr.profilerEnabled = true
	}
}

func UseStorage(s Storage) OptionFn {
	return func(srvr *Server) {
		srvr.storage = s
	}
}

func TLSConfig(cert, pk string) OptionFn {
	certificate, err := tls.LoadX509KeyPair(cert, pk)
	return func(srvr *Server) {
		srvr.tlsConfig = &tls.Config{
			GetCertificate: func(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
				return &certificate, err
			},
		}
	}
}

func (s *Server) Run() {
	log.Println("TRANSFER.SH SERVER STARTED")
	log.Printf("listening for HTTP connections on %s", s.ListenerString)

	if s.TLSListenerString != "" {
		log.Printf("listening for HTTPS connections on %s", s.TLSListenerString)

		if s.forceHTTPs {
			log.Println("force-HTTPS enabled")
		}
	}

	if s.profilerEnabled {
		go func() {
			log.Println("listening for profiler connections on :6060")
			http.ListenAndServe(":6060", nil)
		}()
	}

	if s.rateLimitRequests > 0 {
		log.Printf("limiting requests rate to %d per minute", s.rateLimitRequests)
	}

	var fs http.FileSystem

	fs = &assetfs.AssetFS{
		Asset:    web.Asset,
		AssetDir: web.AssetDir,
		AssetInfo: func(path string) (os.FileInfo, error) {
			return os.Stat(path)
		},
		Prefix: web.Prefix,
	}

	for _, path := range web.AssetNames() {
		bytes, err := web.Asset(path)
		if err != nil {
			log.Panicf("unable to parse: path=%s, err=%s", path, err)
		}

		htmlTemplates.New(stripPrefix(path)).Parse(string(bytes))
		textTemplates.New(stripPrefix(path)).Parse(string(bytes))
	}

	log.Printf("path to temporary files: %s", s.tempPath)
	log.Println("------------------------------------------")

	r := mux.NewRouter()

	staticHandler := http.FileServer(fs)

	mainPage := s.mainPage
	previewPage := s.previewPage

	sendFile := s.sendFile
	sendTarGz := s.sendTarGz
	sendZip := s.sendZip

	postFile := s.postFile
	putFile := s.putFile

	if s.rateLimitRequests > 0 {
		mainPage = ratelimit.Request(ratelimit.IP).Rate(s.rateLimitRequests, 60*time.Second).LimitBy(memory.New())(http.HandlerFunc(mainPage)).ServeHTTP
		previewPage = ratelimit.Request(ratelimit.IP).Rate(s.rateLimitRequests, 60*time.Second).LimitBy(memory.New())(http.HandlerFunc(previewPage)).ServeHTTP

		sendFile = ratelimit.Request(ratelimit.IP).Rate(s.rateLimitRequests, 60*time.Second).LimitBy(memory.New())(http.HandlerFunc(sendFile)).ServeHTTP
		sendTarGz = ratelimit.Request(ratelimit.IP).Rate(s.rateLimitRequests, 60*time.Second).LimitBy(memory.New())(http.HandlerFunc(sendTarGz)).ServeHTTP
		sendZip = ratelimit.Request(ratelimit.IP).Rate(s.rateLimitRequests, 60*time.Second).LimitBy(memory.New())(http.HandlerFunc(sendZip)).ServeHTTP

		postFile = ratelimit.Request(ratelimit.IP).Rate(s.rateLimitRequests, 60*time.Second).LimitBy(memory.New())(http.HandlerFunc(postFile)).ServeHTTP
		putFile = ratelimit.Request(ratelimit.IP).Rate(s.rateLimitRequests, 60*time.Second).LimitBy(memory.New())(http.HandlerFunc(putFile)).ServeHTTP
	}

	r.PathPrefix("/img/").Handler(staticHandler)
	r.PathPrefix("/css/").Handler(staticHandler)
	r.PathPrefix("/js/").Handler(staticHandler)
	r.PathPrefix("/fonts/").Handler(staticHandler)
	r.PathPrefix("/robots.txt").Handler(staticHandler)

	r.HandleFunc("/", mainPage).Methods("GET")
	r.HandleFunc("/{token}/{filename}", previewPage).MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) (match bool) {
		match = false

		// The file will show a preview page when opening the link in browser directly or
		// from external link. If the referer url path and current path are the same it will be
		// downloaded.
		if !acceptsHTML(r.Header) {
			return false
		}

		match = (r.Referer() == "")

		u, err := url.Parse(r.Referer())
		if err != nil {
			log.Fatal(err)
			return
		}

		match = match || (u.Path != r.URL.Path)
		return
	}).Methods("GET")

	r.HandleFunc("/{token}/{filename}", sendFile).Methods("GET")
	r.HandleFunc("/({files:.*}).tar.gz", sendTarGz).Methods("GET")
	r.HandleFunc("/({files:.*}).zip", sendZip).Methods("GET")

	r.HandleFunc("/{filename}", putFile).Methods("PUT")
	r.HandleFunc("/", postFile).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(s.notFoundError)

	h := handlers.PanicHandler(handlers.LogHandler(s.redirectHandler(r), handlers.NewLogOptions(log.Printf, "_default_")), nil)

	srvr := &http.Server{
		Addr:    s.ListenerString,
		Handler: h,
	}

	go func() {
		srvr.ListenAndServe()
	}()

	if s.TLSListenerString != "" {
		go func() {
			s := &http.Server{
				Addr:      s.TLSListenerString,
				Handler:   h,
				TLSConfig: s.tlsConfig,
			}

			if err := s.ListenAndServeTLS("", ""); err != nil {
				panic(err)
			}
		}()
	}

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt)
	signal.Notify(term, syscall.SIGTERM)

	<-term

	log.Println("SERVER STOPPED")
}
