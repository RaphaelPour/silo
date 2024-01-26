package silo

import (
	"errors"
)

var (
	ErrKeyNotExist = errors.New("key doesn't exist")
)

type Driver interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	Delete(key string) error
}
