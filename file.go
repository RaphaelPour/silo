package silo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var (
	ErrWriteFile = errors.New("error writing file")
	ErrReadFile  = errors.New("error reading file")
)

type File struct {
	cache    map[string]any
	filename string
	dirty    bool
}

func NewFile(filename string) Driver {
	return File{
		filename: filename,
		dirty:    true,
		cache:    make(map[string]any),
	}
}

func (f File) Load() error {
	rawJson, err := os.ReadFile(f.filename)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrReadFile, err)
	}

	return json.Unmarshal(rawJson, &f.cache)
}

func (f File) Save() error {
	rawJSON, err := json.Marshal(f.cache)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrWriteFile, err)
	}

	if err := os.WriteFile(f.filename, rawJSON, 0600); err != nil {
		return fmt.Errorf("%w: %w", ErrWriteFile, err)
	}

	return nil
}

func (f File) Get(key string) (any, error) {
	if f.dirty {
		if err := f.Load(); err != nil {
			return nil, err
		}
		f.dirty = false
	}

	value, ok := f.cache[key]
	if !ok {
		return nil, ErrKeyNotExist
	}
	return value, nil
}

func (f File) Set(key string, value any) error {
	f.cache[key] = value
	return f.Save()
}

func (f File) Delete(key string) error {
	delete(f.cache, key)
	return f.Save()
}
