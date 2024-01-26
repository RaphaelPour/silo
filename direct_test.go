package silo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCRUD(t *testing.T) {
	store := NewDirect()
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
