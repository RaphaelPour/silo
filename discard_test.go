package silo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiscard(t *testing.T) {
	store := NewDiscard()
	require.NoError(t, store.Set("a", "b"))

	val, err := store.Get("a")
	require.NoError(t, err)
	require.Nil(t, val)

	require.NoError(t, store.Delete("a"))
}
