package silo

import (
	"errors"
)

var (
	ErrKeyNotExist = errors.New("key doesn't exist")
)

type KeyValueLayer interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	Delete(key string) error
}

type FormatLayer interface {
	Write(data map[string]any) error
	Read() (map[string]any, error)
}

type DataLayer interface {
	Write(data []byte) error
	Read() ([]byte, error)
}
