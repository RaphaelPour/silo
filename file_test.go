package silo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileCRUD(t *testing.T) {
	f, err := os.CreateTemp("", "dump")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	store := NewFile(f.Name())
	require.NoError(t, store.Write([]byte("abc")))

	value, err := store.Read()
	require.NoError(t, err)
	require.Equal(t, []byte("abc"), value)
}

func TestFilePersist(t *testing.T) {
	f, err := os.CreateTemp("", "dump")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	store := NewFile(f.Name())
	require.NoError(t, store.Write([]byte("a")))

	store2 := NewFile(f.Name())
	value, err := store2.Read()
	require.NoError(t, err)
	require.Equal(t, []byte("a"), value)
}

func TestFileBadFile(t *testing.T) {
	store := NewFile("bad/file")
	require.ErrorIs(t, store.Write([]byte("abc")), ErrWriteFile)

	_, err := store.Read()
	require.ErrorIs(t, err, ErrReadFile)
}
