package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/willsmile/s2test/internal/depository"
	"github.com/willsmile/s2test/internal/executor"
)

var (
	// ErrEmptyPath is returned when the path argument is empty
	ErrEmptyPath = errors.New("path argument is empty")
	// ErrReadFile is returned when failed to read a file
	ErrReadFile = errors.New("failed to read a file")
)

type Config interface {
	executor.Plan | depository.Endpoints
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
