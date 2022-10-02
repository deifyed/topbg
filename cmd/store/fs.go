package store

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/afero"
)

func findCurrentImageInDirectory(fs *afero.Afero, dir string, filename string) (image, error) {
	files, err := fs.ReadDir(dir)
	if err != nil {
		return image{}, fmt.Errorf("reading directory: %w", err)
	}

	var relevantFile os.FileInfo

	for _, file := range files {
		if extractFilename(file.Name()) != filename {
			continue
		}

		if relevantFile != nil && relevantFile.ModTime().Before(relevantFile.ModTime()) {
			continue
		}

		relevantFile = file

	}

	if relevantFile == nil {
		return image{}, errors.New("not found")
	}

	content, err := fs.ReadFile(path.Join(dir, relevantFile.Name()))
	if err != nil {
		return image{}, fmt.Errorf("reading file %s: %w", relevantFile.Name(), err)
	}

	return image{
		Image:     bytes.NewReader(content),
		Extension: strings.ReplaceAll(path.Ext(relevantFile.Name()), ".", ""),
	}, nil
}

func extractFilename(fullName string) string {
	parts := strings.Split(fullName, ".")

	return parts[0]
}
