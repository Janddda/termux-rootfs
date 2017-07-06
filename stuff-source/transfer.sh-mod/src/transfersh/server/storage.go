package server

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
)

type Storage interface {
	Basedir() string
	Head(token string, filename string) (contentType string, contentLength uint64, err error)
	IsNotExist(err error) bool
	Get(token string, filename string) (reader io.ReadCloser, contentType string, contentLength uint64, err error)
	Put(token string, filename string, reader io.Reader, contentType string, contentLength uint64) error
	Delete(token string, filename string) error
}

type LocalStorage struct {
	Storage
	basedir string
}

func NewLocalStorage(basedir string) (*LocalStorage, error) {
	return &LocalStorage{basedir: basedir}, nil
}

func (s *LocalStorage) Basedir() string {
	return s.basedir
}

func (s *LocalStorage) Head(token string, filename string) (contentType string, contentLength uint64, err error) {
	path := filepath.Join(s.basedir, token, filename)

	var fi os.FileInfo
	if fi, err = os.Lstat(path); err != nil {
		return
	}

	contentLength = uint64(fi.Size())

	contentType = mime.TypeByExtension(filepath.Ext(filename))

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return
}

func (s *LocalStorage) IsNotExist(err error) bool {
	if err == nil {
		return false
	}

	return os.IsNotExist(err)
}

func (s *LocalStorage) Get(token string, filename string) (reader io.ReadCloser, contentType string, contentLength uint64, err error) {
	path := filepath.Join(s.basedir, token, filename)

	// content type , content length
	if reader, err = os.Open(path); err != nil {
		return
	}

	var fi os.FileInfo
	if fi, err = os.Lstat(path); err != nil {
		return
	}

	contentLength = uint64(fi.Size())

	contentType = mime.TypeByExtension(filepath.Ext(filename))

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return
}

func (s *LocalStorage) Put(token string, filename string, reader io.Reader, contentType string, contentLength uint64) error {
	var f io.WriteCloser
	var err error

	path := filepath.Join(s.basedir, token)

	if err = os.Mkdir(path, 0700); err != nil && !os.IsExist(err) {
		return err
	}

	if f, err = os.OpenFile(filepath.Join(path, filename), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600); err != nil {
		fmt.Printf("%s", err)
		return err
	}

	defer f.Close()

	if _, err = io.Copy(f, reader); err != nil {
		return err
	}

	return nil
}

func (s *LocalStorage) Delete(token string, filename string) error {
	path := filepath.Join(s.basedir, token)

	if err := os.RemoveAll(path); err == nil {
		return errors.New(fmt.Sprintf("storage: file \"/%s/%s\" successfully deleted", token, filename))
	} else {
		return errors.New(fmt.Sprintf("storage: failed to delete \"/%s/%s\" because of %s", token, filename, err.Error()))
	}
}
