package utils

import (
	"errors"
	"log/slog"
	"os"
)

func ReadInputFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("Error reading input file", "path", filePath, "error", err)
		panic(err)
	}
	return string(data)
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	slog.Error("Unexpected error checking file", "fileName", filePath, "error", err)
	return false
}
