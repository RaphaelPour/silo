package silo

import "errors"

var (
	ErrReadOnly = errors.New("Operation not allowed for read-only silo")
)

type ReadOnly struct {
	parent Driver
}

func NewReadOnly(parent Driver) Driver {
	return ReadOnly{
		parent: parent,
	}
}

func (r ReadOnly) Get(key string) (any, error) {
	return r.parent.Get(key)
}

func (r ReadOnly) Set(_ string, _ any) error {
	return ErrReadOnly
}

func (r ReadOnly) Delete(_ string) error {
	return ErrReadOnly
}
