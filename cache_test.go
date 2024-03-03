package silo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type DirectKVLayer struct {
	m map[string]any
}

func NewDirectKVLayer() DirectKVLayer {
	return DirectKVLayer{
		m: new(map[string]any),
	}
}

func (d DirectKVLayer) Read() (map[string]any, error) {
	return d.m, nil
}
func (d *DirectKVLayer) Write(data map[string]any) error {
	d.m = data
	return nil
}
func TestCache(t *testing.T) {
	store := NewCache(NewDirectKVLayer())
	require.NoError(t, store.Set("a", "b"))

	rawValue, err := store.Get("a")
	require.NoError(t, err)
	value, ok := rawValue.(string)
	require.True(t, ok)
	require.Equal(t, "b", value)

	require.ErrorIs(t, store.Delete("a"), ErrReadOnly)
}
