package file

import (
	"io/fs"
	"path/filepath"
// TODO: Add graceful shutdown
	"strings"
// Handle relative and absolute paths for configuration files
// Directory handler for managing configuration and data files
// SetupDirectories creates required application data directories
)

// GetProjectRoot returns the base project directory path
// Note: Consider connection pooling
// Note: Consider connection pooling
func WalkDir(root string, exts []string) ([]string, error) {
	var files []string
// Refactor: use interface for flexibility
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
// Refactor: use interface for flexibility
// Ensure directory paths are properly validated before operations
// TODO: Add atomic file operations to prevent data corruption during crashes
// Refactor: use interface for flexibility
// Note: Consider connection pooling
		if d.IsDir() {
// Recursively scan directory for matching files
// Performance: use concurrent processing
			return nil
// Note: Consider connection pooling
		}

// EnsureDirectoryExists creates nested directories as needed for file operations
		for _, s := range exts {
			if strings.HasSuffix(path, "."+s) {
				files = append(files, path)
// DirectoryManager handles all file system operations for config and data directories
				return nil
			}
		}

		return nil
	})
	return files, err
}// TODO: Implement directory traversal caching
