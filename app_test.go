package main

import (
	"errors"
	"testing"
)

func TestAppRun_ArgsOfValidPath(t *testing.T) {
	args := []string{"./s2test", "-p", "testdata/plan.json"}
	err := New().Run(args)
	if err != nil {
		t.Fatalf("App run with args of valid path, expected none error, got %s", err)
	}
}

func TestAppRun_ArgsWithInvalidPath(t *testing.T) {
	args := []string{"./s2test", "-p", "testdata/invalid_plan.json"}
	err := New().Run(args)
	if !errors.Is(err, ErrReadFile) {
		t.Fatalf("App run with args of invalid path, expected %s, got %s", ErrReadFile, err)
	}
}

func TestAppRun_ArgsWithEmptyPath(t *testing.T) {
	args := []string{"./s2test", "-p", ""}
	err := New().Run(args)
	if !errors.Is(err, ErrEmptyPath) {
		t.Fatalf("App run with args of empty path, expected %s, got %s", ErrEmptyPath, err)
	}
}

func TestAppRun_NoArgs(t *testing.T) {
	args := []string{"./s2test"}
	err := New().Run(args)
	if !errors.Is(err, ErrEmptyPath) {
		t.Fatalf("App run with no args, expected %s, got %s", ErrEmptyPath, err)
	}
}
