package silo

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrWriteFile = errors.New("error writing file")
	ErrReadFile  = errors.New("error reading file")
)

type File struct {
	filename string
}

func NewFile(filename string) DataLayer {
	return &File{
		filename: filename,
	}
}

func (f File) createIfNotExisting() error {
	if _, err := os.Stat(f.filename); os.IsNotExist(err) {
		f, err := os.OpenFile(f.filename, os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	return nil
}

func (f File) Read() ([]byte, error) {
	if err := f.createIfNotExisting(); err != nil {
		return nil, err
	}

	content, err := os.ReadFile(f.filename)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrReadFile, err)
	}
	return content, nil
}

func (f File) Write(data []byte) error {
	if err := f.createIfNotExisting(); err != nil {
		return err
	}

	if err := os.WriteFile(f.filename, data, 0600); err != nil {
		return fmt.Errorf("%w: %w", ErrWriteFile, err)
	}
	return nil
}
