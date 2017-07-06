package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

func (s *Server) lock(token, filename string) error {
	key := path.Join(token, filename)

	if _, ok := s.locks[key]; !ok {
		s.locks[key] = &sync.Mutex{}
	}

	s.locks[key].Lock()

	return nil
}

func (s *Server) unlock(token, filename string) error {
	key := path.Join(token, filename)
	s.locks[key].Unlock()

	return nil
}

type Metadata struct {
	// ContentType is the original uploading content type
	ContentType string
	// Downloads is the actual number of downloads
	Downloads int
	// MaxDownloads contains the maximum numbers of downloads
	MaxDownloads int
	// MaxDate contains the max age of the file
	MaxDate time.Time
}

func metadataForRequest(contentType string, r *http.Request) Metadata {
	metadata := Metadata{
		ContentType:  contentType,
		MaxDate:      time.Now().Add(time.Hour * 24 * 14),
		Downloads:    0,
		MaxDownloads: -1,
	}

	if v := r.Header.Get("Max-Downloads"); v == "" {
	} else if v, err := strconv.Atoi(v); err != nil {
	} else {
		metadata.MaxDownloads = v
	}

	if v := r.Header.Get("Max-Days"); v == "" {
	} else if v, err := strconv.Atoi(v); err != nil {
	} else {
		metadata.MaxDate = time.Now().Add(time.Hour * 24 * time.Duration(v))
	}

	return metadata
}

func (s *Server) checkAgeExpiry(token, filename string) error {
	s.lock(token, filename)
	defer s.unlock(token, filename)

	var metadata Metadata

	path := filepath.Join(s.storage.Basedir(), token)

	_, err := os.Stat(path)
	if err != nil {
		// Don't do anything when token directory is not exist.
		// Get handlers already have this check and will return 404 status.
		return nil
	}

	r, _, _, err := s.storage.Get(token, fmt.Sprintf("%s.metadata", filename))
	if s.storage.IsNotExist(err) {
		log.Printf("metadata check: file \"/%s/%s\" has no metadata !", token, filename)
		return s.storage.Delete(token, filename)
	} else if err != nil {
		return err
	}

	defer r.Close()

	if err := json.NewDecoder(r).Decode(&metadata); err != nil {
		return err
	} else if time.Now().After(metadata.MaxDate) {
		log.Printf("metadata check: max-date of \"/%s/%s\" expired", token, filename)
		return s.storage.Delete(token, filename)
	}

	return nil
}

func (s *Server) checkDownloadsExpiry(token, filename string) error {
	s.lock(token, filename)
	defer s.unlock(token, filename)

	var metadata Metadata

	r, _, _, err := s.storage.Get(token, fmt.Sprintf("%s.metadata", filename))
	if s.storage.IsNotExist(err) {
		log.Printf("metadata check: file \"/%s/%s\" has no metadata !", token, filename)
		return s.storage.Delete(token, filename)
	} else if err != nil {
		return err
	}

	defer r.Close()

	if err := json.NewDecoder(r).Decode(&metadata); err != nil {
		return err
	} else if metadata.MaxDownloads > 0 && metadata.Downloads >= metadata.MaxDownloads {
		log.Printf("metadata check: max-downloads of \"/%s/%s\" expired", token, filename)
		return s.storage.Delete(token, filename)
	} else {
		// update number of downloads
		metadata.Downloads++

		buffer := &bytes.Buffer{}
		if err := json.NewEncoder(buffer).Encode(metadata); err != nil {
			return errors.New("could not encode metadata")
		} else if err := s.storage.Put(token, fmt.Sprintf("%s.metadata", filename), buffer, "text/json", uint64(buffer.Len())); err != nil {
			return errors.New("could not save metadata")
		}

		if metadata.MaxDownloads > 0 && metadata.Downloads >= metadata.MaxDownloads {
			log.Printf("metadata check: max-downloads of \"/%s/%s\" expired", token, filename)
			return s.storage.Delete(token, filename)
		}
	}

	return nil
}
