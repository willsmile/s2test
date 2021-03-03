package main

import (
	"errors"
	"testing"
)

func TestLoadStore(t *testing.T) {
	_, err := LoadStore("./testdata/api.json")
	if err != nil {
		t.Fatalf("LoadStore(\"./testdata/api.json\"), expected none error, got %s", err)
	}
}

func TestLoadStore_emptyPath(t *testing.T) {
	_, err := LoadStore("")
	if !errors.Is(err, ErrEmptyPath) {
		t.Fatalf("LoadStore(\"\"), expected %s, got %s", ErrEmptyPath, err)
	}
}

func TestLoadStore_invalidPath(t *testing.T) {
	_, err := LoadStore("./testdata/invalid_api.json")
	if !errors.Is(err, ErrReadFile) {
		t.Fatalf("LoadStore(\"./testdata/invalid_api.json\"), expected %s, got %s", ErrReadFile, err)
	}
}
