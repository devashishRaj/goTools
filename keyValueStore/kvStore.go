package keyValueStore

import (
	"encoding/gob"
	"errors"
	"io/fs"
	"os"
)

type Store struct {
	path string
	data map[string]string
}

func OpenStore(path string) (*Store, error) {
	s := &Store{path: path, data: map[string]string{}}
	f, err := os.Open(path)
	// fs.ErrNotExist is an error that occurs , if there's no file at given path
	// thus store is new one, and we would just return it
	if errors.Is(err, fs.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	err = gob.NewDecoder(f).Decode(&s.data)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Store) Set(k, v string) {
	s.data[k] = v
}

func (s *Store) Get(k string) (string, bool) {
	v, ok := s.data[k]
	return v, ok
}

func (s *Store) Save() error {
	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	return gob.NewEncoder(f).Encode(s.data)
}
