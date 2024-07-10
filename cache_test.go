package silo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestLayer struct {
	data map[string]any
}

func NewTestLayer() FormatLayer {
	return &TestLayer{
		data: make(map[string]any),
	}
}

func (t TestLayer) Read() (map[string]any, error) {
	return t.data, nil
}

func (t *TestLayer) Write(data map[string]any) error {
	t.data = data
	return nil
}

func TestCache(t *testing.T) {
	store := NewCache(NewTestLayer())

	t.Run("get fails for not existing key", func(t *testing.T) {
		_, err := store.Get("a")
		require.ErrorIs(t, ErrKeyNotExist, err)
	})

	t.Run("set key", func(t *testing.T) {
		require.NoError(t, store.Set("a", "b"))
	})

	t.Run("get key", func(t *testing.T) {
		rawValue, err := store.Get("a")
		require.NoError(t, err)
		value, ok := rawValue.(string)
		require.True(t, ok)
		require.Equal(t, "b", value)
	})

	t.Run("delete key", func(t *testing.T) {
		require.NoError(t, store.Delete("a"))

		_, err := store.Get("a")
		require.ErrorIs(t, ErrKeyNotExist, err)
	})
}

type BadLayer struct{}

func (b BadLayer) Read() (map[string]any, error) {
	return nil, errors.New("fatal disk error")
}

func (b BadLayer) Write(_ map[string]any) error {
	return errors.New("fatal disk error")
}

func TestCacheBadParent(t *testing.T) {
	t.Run("get fails if parent load fails", func(t *testing.T) {
		store := NewCache(BadLayer{})
		_, err := store.Get("some-key")
		require.Error(t, err)
	})

	t.Run("set fails if parent load fails", func(t *testing.T) {
		store := NewCache(BadLayer{})
		require.Error(t, store.Set("some-key", "some-value"))
	})

	t.Run("delete fails if parent load fails", func(t *testing.T) {
		store := NewCache(BadLayer{})
		require.Error(t, store.Delete("some-key"))
	})
}
