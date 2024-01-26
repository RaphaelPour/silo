package silo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var (
	ErrKeyNotExist = errors.New("key doesn't exist")
)

type Driver interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	Delete(key string) error
}

type Direct struct {
	store map[string]any
}

func NewDirect() Driver {
	return Direct{}
}

func (d Direct) Get(key string) (any, error) {
	value, ok := d.store[key]
	if !ok {
		return nil, ErrKeyNotExist
	}
	return value, nil
}

func (d Direct) Set(key string, value any) error {
	d.store[key] = value
	return nil
}

func (d Direct) Delete(key string) error {
	delete(d.store, key)
	return nil
}

type File struct {
	cache    map[string]any
	filename string
	dirty    bool
}

func NewFile(filename string) Driver {
	return File{
		filename: filename,
		dirty:    true,
	}
}

func (f File) Load() error {
	rawJson, err := os.ReadFile(f.filename)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return json.Unmarshal(rawJson, &f.cache)
}

func (f File) Save() error {
	rawJSON, err := json.Marshal(f.cache)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}
	return os.WriteFile(f.filename, rawJSON, 0600)
}

func (f File) Get(key string) (any, error) {
	if f.dirty {
		if err := f.Load(); err != nil {
			return nil, err
		}
		f.dirty = false
	}
	return f.cache[key], nil
}

func (f File) Set(key string, value any) error {
	f.cache[key] = value
	return f.Save()
}

func (f File) Delete(key string) error {
	delete(f.cache, key)
	return f.Save()
}
