package wm

import (
	"fmt"
	"io"
	"path"

	"github.com/deifyed/topbg/pkg/config"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

const TemporaryFilename = "current-topbg"

func SetBackground(fs *afero.Afero, imageType string, image io.Reader) error {
	imagePath := getImagePath(imageType)

	err := fs.WriteReader(imagePath, image)
	if err != nil {
		return fmt.Errorf("writing image: %w", err)
	}

	err = swayset(imagePath)
	if err != nil {
		return fmt.Errorf("setting background: %w", err)
	}

	return nil
}

func getImagePath(imageType string) string {
	return path.Join(
		viper.GetString(config.TemporaryImageDir),
		fmt.Sprintf("%s.%s", TemporaryFilename, imageType),
	)
}
