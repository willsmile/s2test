package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/willsmile/s2test/internal/executor"
	"github.com/willsmile/s2test/internal/storage"
)

func TestLoadJSON_Store(t *testing.T) {
	endpoints := storage.Endpoints{}
	err := LoadJSON("../../testdata/api.json", &endpoints)
	if err != nil {
		t.Fatalf("LoadJSON(\"../../testdata/api.json\", &endpoints), expected none error, got %s", err)
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
	endpoints := storage.Endpoints{}
	err := LoadJSON("", &endpoints)
	if !errors.Is(err, ErrEmptyPath) {
		t.Fatalf("LoadJSON(\"\", &endpoints), expected %s, got %s", ErrEmptyPath, err)
	}
}

func TestLoadJSON_WithInvalidPath(t *testing.T) {
	endpoints := storage.Endpoints{}
	err := LoadJSON("../../testdata/invalid_api.json", &endpoints)
	if !errors.Is(err, ErrReadFile) {
		t.Fatalf("LoadJSON(\"../../testdata/invalid_api.json\", &endpoints), expected %s, got %s", ErrReadFile, err)
	}
}
