package wm

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"path"

	"github.com/spf13/afero"
)

func SetBackground(fs *afero.Afero, imageType string, image io.Reader) error {
        imagePath := path.Join(fs.GetTempDir(""), fmt.Sprintf("current-topbg.%s", imageType))

        err := fs.WriteReader(imagePath, image)
        if err != nil {
                return fmt.Errorf("writing image: %w", err)
        }

        swayset(imagePath)

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
