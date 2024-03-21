package silo

import (
	"encoding/json"
)

type Json struct {
	parent DataLayer
}

func NewJson(parent DataLayer) FormatLayer {
	return Json{
		parent: parent,
	}
}

func (j Json) Read() (map[string]any, error) {
	bytes, err := j.parent.Read()
	if err != nil {
		return nil, err
	}

	// initialize map if daa is empty, e.g  empty file was just created
	if len(bytes) == 0 {
		return map[string]any{}, nil
	}

	var m map[string]any
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (j Json) Write(jsonData map[string]any) error {
	// if map is empty, marshaller will write "null" which
	// can't be converted back into a map
	if jsonData == nil {
		return j.parent.Write([]byte("{}"))
	}
	bytes, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}
	return j.parent.Write(bytes)
}
