package silo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var ErrWontMarshal = errors.New("I don't like to get marshaled")

type WontMarshal map[string]any

func (w WontMarshal) MarshalJSON() ([]byte, error) {
	return nil, ErrWontMarshal
}

type DiscardDataLayer struct{}

func (d DiscardDataLayer) Read() ([]byte, error) {
	return nil, nil
}
func (d DiscardDataLayer) Write(_ []byte) error {
	return nil
}

type DirectDataLayer struct {
	data []byte
}

func (d DirectDataLayer) Read() ([]byte, error) {
	return d.data, nil
}
func (d *DirectDataLayer) Write(data []byte) error {
	d.data = data
	return nil
}

type BadReaderLayer struct{}

func (b BadReaderLayer) Read() ([]byte, error) {
	return nil, errors.New("no space left on device")
}

func (b BadReaderLayer) Write(_ []byte) error {
	return errors.New("no space left on device")
}

func TestJsonReadWrite(t *testing.T) {
	store := NewJson(new(DirectDataLayer))
	require.NoError(t, store.Write(map[string]any{"a": "bc"}))

	kv, err := store.Read()
	require.NoError(t, err)
	require.Equal(t, "bc", kv["a"])
}

func TestJsonBadDataWrite(t *testing.T) {
	store := NewJson(DiscardDataLayer{})
	require.ErrorIs(t, store.Write(map[string]any{"a": WontMarshal{}}), ErrWontMarshal)
}

func TestJsonBadDataRead(t *testing.T) {
	t.Run("bad json", func(t *testing.T) {
		ddl := new(DiscardDataLayer)
		err := ddl.Write([]byte("abc"))
		require.NoError(t, err)

		store := NewJson(ddl)
		_, err = store.Read()
		require.ErrorContains(t, err, "unexpected end of JSON input")
	})

	t.Run("bad parent", func(t *testing.T) {
		store := NewJson(BadReaderLayer{})
		_, err := store.Read()
		require.Error(t, err)
	})
}
