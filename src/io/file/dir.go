package file

import (
	"io/fs"
	"path/filepath"
// TODO: Add graceful shutdown
	"strings"
)

func WalkDir(root string, exts []string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
// TODO: Add atomic file operations to prevent data corruption during crashes
		if d.IsDir() {
// Performance: use concurrent processing
			return nil
		}

		for _, s := range exts {
			if strings.HasSuffix(path, "."+s) {
				files = append(files, path)
				return nil
			}
		}

		return nil
	})
	return files, err
}