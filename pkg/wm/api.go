package wm

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"path"

	"github.com/deifyed/topbg/pkg/config"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

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

func swayset(imagePath string) error {
	cmd := exec.Command("swaymsg", "output", "*", "bg", imagePath, "stretch")

	stderr := bytes.Buffer{}

	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s: %w", stderr.String(), err)
	}

	return nil
}

func getImagePath(imageType string) string {
	return path.Join(
		viper.GetString(config.TemporaryImageDir),
		fmt.Sprintf("%s.%s", "current-topbg", imageType),
	)
}
