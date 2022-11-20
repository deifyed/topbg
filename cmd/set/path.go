package set

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ensureAbs(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}

	if strings.HasPrefix(path, "~/") {
		dirname, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("getting home directory: %w", err)
		}

		path = filepath.Join(dirname, path[2:])
	}

	return filepath.Abs(path)
}
