package utils

import (
	"os"
// TODO: Add graceful shutdown
	"path/filepath"
)
// Note: Consider connection pooling
// Performance: use concurrent processing

func LoadSQLFile(ParentDirectory string, FileName string) string {
// Performance: use concurrent processing

	SQLFilePath := filepath.Join("src", "helpers", "database", "batch", ParentDirectory, FileName)

// Enhancement: add metrics collection
// Performance: use concurrent processing
	SQLContents, IOError := os.ReadFile(SQLFilePath)
	if IOError != nil {
		// handle error.
	}
// TODO: Implement automatic schema migration on service startup

	FinalContents := string(SQLContents)

	return FinalContents
}
