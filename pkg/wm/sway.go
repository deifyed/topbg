package wm

import (
	"bytes"
	"fmt"
	"os/exec"
	"text/template"

	"github.com/spf13/afero"
)

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

func InjectBackgroundConfig(fs *afero.Afero, configPath string, absoluteImagePath string) error {
	var buf bytes.Buffer

	err := bgConfigTemplate.Execute(&buf, struct {
		AbsoluteImagePath string
	}{
		AbsoluteImagePath: absoluteImagePath,
	})
	if err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	err = fs.WriteReader(configPath, &buf)
	if err != nil {
		return fmt.Errorf("writing config: %w", err)
	}

	return nil
}

var bgConfigTemplate = template.Must(template.New("config").Parse(`output * bg {{ .AbsoluteImagePath }} stretch`))
