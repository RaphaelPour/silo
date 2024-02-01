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

type TestStruct struct {
	number int
	text   string
	list   []int
}

func TestDirectStruct(t *testing.T) {
	input := TestStruct{
		1337, "I am here", []int{3, 2, 1},
	}

	store := NewDirect()
	require.NoError(t, store.Set("key", input))

	rawValue, err := store.Get("key")
	require.NoError(t, err)
	actual, ok := rawValue.(TestStruct)
	require.True(t, ok)
	require.Equal(t, input, actual)
}
