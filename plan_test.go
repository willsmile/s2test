package main

import (
	"errors"
	"testing"
)

func TestLoadPlan(t *testing.T) {
	_, err := LoadPlan("./testdata/plan.json")
	if err != nil {
		t.Fatalf("LoadPlan(\"./testdata/plan.json\"), expected none error, got %s", err)
	}
}

func TestLoadPlan_emptyPath(t *testing.T) {
	_, err := LoadPlan("")
	if !errors.Is(err, ErrEmptyPath) {
		t.Fatalf("LoadPlan(\"\"), expected %s, got %s", ErrEmptyPath, err)
	}
}

func TestLoadPlan_invalidPath(t *testing.T) {
	_, err := LoadPlan("./testdata/invalid_plan.json")
	if !errors.Is(err, ErrReadFile) {
		t.Fatalf("LoadPlan(\"./testdata/invalid_plan.json\"), expected %s, got %s", ErrReadFile, err)
	}
}
