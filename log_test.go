package silo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogCRUD(t *testing.T) {
	logFile, err := os.CreateTemp("", "log")
	require.NoError(t, err)
	defer os.Remove(logFile.Name())

	store, err := NewLog(NewDirect(), logFile.Name())
	require.NoError(t, err)
	require.NoError(t, store.Set("a", "b"))

	rawValue, err := store.Get("a")
	require.NoError(t, err)
	value, ok := rawValue.(string)
	require.True(t, ok)
	require.Equal(t, "b", value)

	require.NoError(t, store.Delete("a"))

	_, err = store.Get("a")
	require.ErrorIs(t, err, ErrKeyNotExist)

	// check if log file has proper values
	rawContent, err := os.ReadFile(logFile.Name())
	require.NoError(t, err)
	require.Regexp(t, "SET.*a=b.*err=<nil>", string(rawContent))
	require.Regexp(t, "GET.*a=b.*err=<nil>", string(rawContent))
	require.Regexp(t, "DELETE.*a.*err=<nil>", string(rawContent))
}

func TestLogBadFile(t *testing.T) {
	_, err := NewLog(NewDirect(), "/in/val/id")
	require.ErrorIs(t, err, ErrOpenLogFile)
}
