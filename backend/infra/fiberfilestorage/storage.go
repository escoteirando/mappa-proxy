package fiberfilestorage

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

type (
	FiberFileStorage struct {
		logger    *log.Logger
		config    *Config
		mux       sync.RWMutex
		done      chan struct{}
		filenamer func(string) string
	}
	Config struct {
		BasePath   string
		GCInterval time.Duration
	}
)

func New(config *Config) *FiberFileStorage {
	if stat, err := os.Stat(config.BasePath); err != nil || !stat.IsDir() {
		err = os.MkdirAll(config.BasePath, 0755)
		if err != nil {
			panic(err)
		}
	}

	store := &FiberFileStorage{
		config:    config,
		logger:    log.New(os.Stdout, "FileStorage: ", log.LstdFlags),
		filenamer: str_filename,
	}
	store.logger.Printf("Using %s as storage", config.BasePath)
	go store.gc()
	return store
}

func (s *FiberFileStorage) Get(key string) ([]byte, error) {
	content, validUntil, err := s.readData(key)
	if err != nil {
		return nil, err
	}
	if validUntil.Before(time.Now()) {
		return nil, os.ErrNotExist
	}
	return content, nil
}

func (s *FiberFileStorage) Set(key string, val []byte, exp time.Duration) error {
	var validUntil time.Time
	if exp > 0 {
		validUntil = time.Now().Add(exp)
	} else {
		validUntil = time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
	}
	return s.writeData(key, val, validUntil)
}

func (s *FiberFileStorage) Delete(key string) error {
	return s.deleteData(key)
}

func (s *FiberFileStorage) Reset() error {
	entries, err := os.ReadDir(s.config.BasePath)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		err := os.Remove(path.Join(s.config.BasePath, entry.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *FiberFileStorage) Close() error {
	s.done <- struct{}{}
	return nil
}

func md5_filename(filename string) string {
	hasher := md5.New()
	hasher.Write([]byte(filename))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

func str_filename(filename string) string {
	return strings.ReplaceAll(filename, "/", "_")
}

func (s *FiberFileStorage) fileName(key string) string {
	fn := s.filenamer(key)
	return path.Join(s.config.BasePath, fn)
}

func (s *FiberFileStorage) readData(key string) (content []byte, validUntil time.Time, err error) {
	filename := s.fileName(key)
	if stat, err := os.Stat(filename); err != nil || stat.IsDir() {
		return nil, time.Time{}, err
	}
	fullContent, err := ioutil.ReadFile(filename)
	vu := fullContent[:15]
	if err = validUntil.UnmarshalBinary(vu); err != nil {
		return nil, time.Time{}, err
	}
	content = fullContent[15:]
	return content, validUntil, nil
}

func (s *FiberFileStorage) writeData(key string, content []byte, validUntil time.Time) (err error) {
	filename := s.fileName(key)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	vu, err := validUntil.MarshalBinary()
	if err != nil {
		return err
	}
	_, err = file.Write(vu)
	if err != nil {
		return err
	}
	_, err = file.Write(content)
	return err

}

func (s *FiberFileStorage) deleteData(key string) (err error) {
	filename := s.fileName(key)
	return os.Remove(filename)
}

func (s *FiberFileStorage) isExpired(filename string) (expired bool, err error) {
	if stat, err := os.Stat(filename); err != nil || stat.IsDir() {
		return false, err
	}
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()
	vu := make([]byte, 15)
	if _, err = file.Read(vu); err != nil {
		return false, err
	}

	var validUntil time.Time
	if err = validUntil.UnmarshalBinary(vu); err != nil {
		return false, err
	}

	return validUntil.Before(time.Now()), nil
}

func (s *FiberFileStorage) gc() {
	s.logger.Printf("Starting GC with interval %s", s.config.GCInterval)
	ticker := time.NewTicker(s.config.GCInterval)
	defer ticker.Stop()
	for {
		select {
		case <-s.done:
			return
		case <-ticker.C:
			s.mux.Lock()
			if entries, err := os.ReadDir(s.config.BasePath); err == nil {
				for _, entry := range entries {
					if entry.IsDir() {
						continue
					}
					filename := path.Join(s.config.BasePath, entry.Name())
					if isExpired, err := s.isExpired(filename); err == nil && isExpired {
						s.logger.Printf("Removed expired file %s - %v", filename, os.Remove(filename))
					}
				}
			}
			s.mux.Unlock()
		}
	}
}
