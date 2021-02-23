package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Store uses for providing a struct for storing API spec information
type Store map[string]spec

// spec uses for providing a struct for information of a single API
type spec struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}

// LoadStore uses for loading store from a JSON file
func LoadStore(path string) Store {
	var s Store

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
