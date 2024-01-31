package silo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLock(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	lock := make(chan struct{}, 1)
	lock <- struct{}{}

	store := Lock{
		parent: NewDirect(),
		ctx:    ctx,
		mutex:  lock,
	}
	require.ErrorIs(t, store.Set("a", "b"), ErrLocking)
	_, err := store.Get("a")
	require.ErrorIs(t, err, ErrLocking)
	require.ErrorIs(t, store.Delete("a"), ErrLocking)
}

func TestLockCRUD(t *testing.T) {
	store := NewLock(context.Background(), NewDirect())
	require.NoError(t, store.Set("a", "b"))

	rawValue, err := store.Get("a")
	require.NoError(t, err)
	value, ok := rawValue.(string)
	require.True(t, ok)
	require.Equal(t, "b", value)

	require.NoError(t, store.Delete("a"))

	_, err = store.Get("a")
	require.ErrorIs(t, ErrKeyNotExist, err)
}
