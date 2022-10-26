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

func injectBackgroundConfig(fs *afero.Afero, configPath string, absoluteImagePath string) error {
	content, err := fs.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("reading: %w", err)
	}

	t, err := template.New("").Parse(bgConfigTemplate)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	var buf bytes.Buffer

	err = t.Execute(&buf, struct {
		ConfigStartPrefix string
		ConfigEndSuffix   string
		AbsoluteImagePath string
	}{
		ConfigStartPrefix: bgConfigPrefix,
		ConfigEndSuffix:   bgConfigSuffix,
		AbsoluteImagePath: absoluteImagePath,
	})
	if err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	newCfgContent := append(content, buf.Bytes()...)

	err = fs.WriteReader(configPath, bytes.NewReader(newCfgContent))
	if err != nil {
		return fmt.Errorf("writing: %w", err)
	}

	return nil
}

const (
	bgConfigPrefix   = "### TOPBG START INJECTED CONFIG ###"
	bgConfigSuffix   = "### TOPBG END INJECTED CONFIG ###"
	bgConfigTemplate = `
{{ .ConfigStartPrefix }}
output * bg {{ .AbsoluteImagePath }} stretch
{{ .ConfigEndSuffix }}
`
)
