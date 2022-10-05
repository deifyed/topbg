package images

import (
	"fmt"
	"path"
	"sort"

	"github.com/deifyed/topbg/pkg/config"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

func ListPaths(fs *afero.Afero) ([]string, error) {
	imageDir := viper.GetString(config.PermanentImageDir)

	files, err := fs.ReadDir(imageDir)
	if err != nil {
		return nil, fmt.Errorf("listing files in directory: %w", err)
	}

	imagePaths := make([]string, len(files))

	for index, file := range files {
		imagePaths[index] = path.Join(imageDir, file.Name())
	}

	return sort.StringSlice(imagePaths), nil
}
