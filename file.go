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

func (f File) Read() ([]byte, error) {
	content, err := os.ReadFile(f.filename)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrReadFile, err)
	}
	return content, nil
}

func (f File) Write(data []byte) error {
	if err := os.WriteFile(f.filename, data, 0600); err != nil {
		return fmt.Errorf("%w: %w", ErrWriteFile, err)
	}
	return nil
}
