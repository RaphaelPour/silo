package silo

import (
	"encoding/json"
	"fmt"
	"os"
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
