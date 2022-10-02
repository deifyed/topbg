package cmd

import (
	"path"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestFindCurrentImageInDirectory(t *testing.T) {
	testCases := []struct {
		name              string
		withExistingPaths []string
		withDirectory     string
		withFilename      string
		expectExtension   string
	}{
		{
			name:              "Should return the correct extension",
			withExistingPaths: []string{"/tmp/somefile", "/tmp/correctfile.jpg", "/tmp/someotherfile"},
			withDirectory:     "/tmp",
			withFilename:      "correctfile",
			expectExtension:   "jpg",
		},
		{
			name:              "Should return the correct extension with two relevant files with different timestamps",
			withExistingPaths: []string{"/tmp/somefile", "/tmp/correctfile.jpg", "/tmp/correctfile.png", "/tmp/someotherfile"},
			withDirectory:     "/tmp",
			withFilename:      "correctfile",
			expectExtension:   "png",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			fs := &afero.Afero{Fs: afero.NewMemMapFs()}

			for _, item := range tc.withExistingPaths {
				fs.MkdirAll(path.Dir(item), 0o700)
				fs.WriteReader(item, strings.NewReader(""))
			}

			img, err := findCurrentImageInDirectory(fs, tc.withDirectory, tc.withFilename)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectExtension, img.Extension)
		})
	}
}
