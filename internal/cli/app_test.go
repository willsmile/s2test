package cli

import (
	"errors"
	"testing"

	"github.com/willsmile/s2test/internal/config"
)

func TestAppRun_ArgsOfValidPath(t *testing.T) {
	args := []string{"./s2test", "-p", "../../testdata/plan.json"}
	err := New().Run(args)
	if err != nil {
		t.Fatalf("App run with args of valid path, expected none error, got %s", err)
	}
}

func TestAppRun_ArgsWithInvalidPath(t *testing.T) {
	args := []string{"./s2test", "-p", "../../testdata/invalid_plan.json"}
	err := New().Run(args)
	if !errors.Is(err, config.ErrReadFile) {
		t.Fatalf("App run with args of invalid path, expected %s, got %s", config.ErrReadFile, err)
	}
}

func TestAppRun_ArgsWithEmptyPath(t *testing.T) {
	args := []string{"./s2test", "-p", ""}
	err := New().Run(args)
	if !errors.Is(err, config.ErrEmptyPath) {
		t.Fatalf("App run with args of empty path, expected %s, got %s", config.ErrEmptyPath, err)
	}
}

func TestAppRun_NoArgs(t *testing.T) {
	args := []string{"./s2test"}
	err := New().Run(args)
	if !errors.Is(err, config.ErrEmptyPath) {
		t.Fatalf("App run with no args, expected %s, got %s", config.ErrEmptyPath, err)
	}
}
