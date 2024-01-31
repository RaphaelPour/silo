package silo

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrLocking = errors.New("error locking")
)

type Lock struct {
	parent Driver
	mutex  chan struct{}
	ctx    context.Context
}

func NewLock(ctx context.Context, parent Driver) Driver {
	return Lock{
		parent: parent,
		mutex:  make(chan struct{}, 1),
		ctx:    ctx,
	}
}

func (l Lock) Get(key string) (any, error) {
	select {
	case l.mutex <- struct{}{}:
		defer func() { <-l.mutex }()
		return l.parent.Get(key)
	case <-l.ctx.Done():
		return nil, fmt.Errorf("%w: %w", ErrLocking, l.ctx.Err())
	}
}

func (l Lock) Set(key string, value any) error {
	select {
	case l.mutex <- struct{}{}:
		defer func() { <-l.mutex }()
		return l.parent.Set(key, value)
	case <-l.ctx.Done():
		return fmt.Errorf("%w: %w", ErrLocking, l.ctx.Err())
	}
}

func (l Lock) Delete(key string) error {
	select {
	case l.mutex <- struct{}{}:
		defer func() { <-l.mutex }()
		return l.parent.Delete(key)
	case <-l.ctx.Done():
		return fmt.Errorf("%w: %w", ErrLocking, l.ctx.Err())
	}
}
