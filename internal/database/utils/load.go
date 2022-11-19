package utils

import (
	"os"
	"path/filepath"
)

func LoadSQLFile(ParentDirectory string, FileName string) string {

	SQLFilePath := filepath.Join("static", "sql", ParentDirectory, FileName)

	SQLContents, IOError := os.ReadFile(SQLFilePath)
	if IOError != nil {
		// handle error.
	}

	FinalContents := string(SQLContents)

	return FinalContents
}
