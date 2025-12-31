package utils

import (
	"embed"
	"log/slog"
)

func ReadInputFile(fs embed.FS, fileName string) string {
	data, err := fs.ReadFile(fileName)
	if err != nil {
		slog.Error("Error reading input file", "file", fileName, "error", err)
		panic(err)
	}
	return string(data)
}
