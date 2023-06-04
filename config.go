package main

import (
	"encoding/json"
	"os"
)

type Config interface {
	Plan | Endpoints
}

func LoadJSON[T Config](path string, t *T) error {
	if path == "" {
		return ErrEmptyPath
	}

	src, error := os.ReadFile(path)
	if error != nil {
		return ErrReadFile
	}

	json.Unmarshal(src, &t)

	return nil
}
