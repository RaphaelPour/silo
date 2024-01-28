package silo

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileCRUD(t *testing.T) {
	f, err := os.CreateTemp("", "dump")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	store := NewFile(f.Name())
	require.NoError(t, store.Set("a", "b"))

	rawValue, err := store.Get("a")
	require.NoError(t, err)
	value, ok := rawValue.(string)
	require.True(t, ok)
	require.Equal(t, "b", value)

	require.NoError(t, store.Delete("a"))

	_, err = store.Get("a")
	require.ErrorIs(t, err, ErrKeyNotExist)
}

type WontMarshal int

func (w WontMarshal) MarshalJSON() ([]byte, error) {
	return nil, errors.New("I don't like to get marshaled")
}

func TestFilePersist(t *testing.T) {
	f, err := os.CreateTemp("", "dump")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	store := NewFile(f.Name())
	require.NoError(t, store.Set("a", "b"))

	store2 := NewFile(f.Name())
	rawValue, err := store2.Get("a")
	require.NoError(t, err)
	value, ok := rawValue.(string)
	require.True(t, ok)
	require.Equal(t, "b", value)
}

func TestFileBadFile(t *testing.T) {
	store := NewFile("bad/file")
	require.ErrorIs(t, store.Set("a", "b"), ErrWriteFile)

	_, err := store.Get("a")
	require.ErrorIs(t, err, ErrReadFile)
}

func TestFileBadData(t *testing.T) {
	f, err := os.CreateTemp("", "dump")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	store := NewFile(f.Name())
	require.ErrorIs(t, store.Set("a", WontMarshal(1)), ErrWriteFile)

	fmt.Println(store.Get("a"))
}
