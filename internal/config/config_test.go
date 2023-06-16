package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/willsmile/s2test/internal/connector"
	"github.com/willsmile/s2test/internal/executor"
)

func TestLoadJSON_Store(t *testing.T) {
	store := connector.Endpoints{}
	err := LoadJSON("../../testdata/api.json", &store)
	if err != nil {
		t.Fatalf("LoadJSON(\"../../testdata/api.json\", &store), expected none error, got %s", err)
	}
}

func TestLoadJSON_Plan(t *testing.T) {
	plan := executor.Plan{}
	err := LoadJSON("../../testdata/plan.json", &plan)
	fmt.Println(plan)
	if err != nil {
		t.Fatalf("LoadJSON(\"../../testdata/api.json\", &plan), expected none error, got %s", err)
	}
}

func TestLoadJSON_WithEmptyPath(t *testing.T) {
	store := connector.Endpoints{}
	err := LoadJSON("", &store)
	if !errors.Is(err, ErrEmptyPath) {
		t.Fatalf("LoadJSON(\"\", &store), expected %s, got %s", ErrEmptyPath, err)
	}
}

func TestLoadJSON_WithInvalidPath(t *testing.T) {
	store := connector.Endpoints{}
	err := LoadJSON("../../testdata/invalid_api.json", &store)
	if !errors.Is(err, ErrReadFile) {
		t.Fatalf("LoadJSON(\"../../testdata/invalid_api.json\", &store), expected %s, got %s", ErrReadFile, err)
	}
}