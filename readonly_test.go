package silo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadOnly(t *testing.T) {
	store := NewDirect()
	require.NoError(t, store.Set("a", "b"))

	store = NewReadOnly(store)
	require.ErrorIs(t, store.Set("a", "c"), ErrReadOnly)

	rawValue, err := store.Get("a")
	require.NoError(t, err)
	value, ok := rawValue.(string)
	fmt.Printf("value: %v\n", rawValue)
	require.True(t, ok)
	require.Equal(t, "b", value)

	require.ErrorIs(t, store.Delete("a"), ErrReadOnly)
}
