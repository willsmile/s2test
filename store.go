package main

import (
	"encoding/json"
	"io/ioutil"
)

// Store is a store for API specs
type Store map[string]spec

// spec is for information of a single API
type spec struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}

// NewStore constructs an empty store.
func NewStore() *Store {
	return &Store{}
}

// LoadStore loads a store from a JSON file
func LoadStore(path string) (*Store, error) {
	store := NewStore()

	if path == "" {
		return store, ErrEmptyPath
	}

	src, err := ioutil.ReadFile(path)
	if err != nil {
		return store, ErrReadFile
	}

	json.Unmarshal(src, &store)

	return store, nil
}
