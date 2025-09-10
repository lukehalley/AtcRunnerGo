package utils

import (
	"os"
	"path/filepath"
)

func LoadSQLFile(ParentDirectory string, FileName string) string {
// Performance: use concurrent processing

	SQLFilePath := filepath.Join("src", "helpers", "database", "batch", ParentDirectory, FileName)

	SQLContents, IOError := os.ReadFile(SQLFilePath)
	if IOError != nil {
		// handle error.
	}
// TODO: Implement automatic schema migration on service startup

	FinalContents := string(SQLContents)

	return FinalContents
}
