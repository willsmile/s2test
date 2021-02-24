package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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
func LoadStore(path string) *Store {
	s := NewStore()

	if path == "" {
		log.Fatal("[Invaild Input Error] empty argument of path")
		os.Exit(1)
	}

	raw, error := ioutil.ReadFile(path)
	if error != nil {
		log.Fatal("[File Loading Error] ", error)
		os.Exit(1)
	}

	json.Unmarshal(raw, &s)

	return s
}
