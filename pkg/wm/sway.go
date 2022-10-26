package wm

import (
	"bytes"
	"errors"
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

	cleanedConfig, err := deleteOldInjectedBackgroundConfig(content)
	if err != nil {
		return fmt.Errorf("deleting old injected background config: %w", err)
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
		ConfigStartPrefix: bgConfigStartDelimiter,
		ConfigEndSuffix:   bgConfigEndDelimiter,
		AbsoluteImagePath: absoluteImagePath,
	})
	if err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	cleanedConfig = append(cleanedConfig, buf.Bytes()...)

	err = fs.WriteReader(configPath, bytes.NewReader(cleanedConfig))
	if err != nil {
		return fmt.Errorf("writing: %w", err)
	}

	return nil
}

func deleteOldInjectedBackgroundConfig(original []byte) ([]byte, error) {
	var (
		ignore        bool
		foundStart    bool
		foundEnd      bool
		cleanedConfig = make([]byte, 0)
	)

	for _, line := range bytes.Split(original, []byte("\n")) {
		if bytes.HasPrefix(line, []byte(bgConfigStartDelimiter)) {
			foundStart = true
			ignore = true
		}

		if bytes.HasPrefix(line, []byte(bgConfigEndDelimiter)) {
			foundEnd = true
			ignore = false

			continue
		}

		if ignore {
			continue
		}

		line = append(line, []byte("\n")...)
		cleanedConfig = append(cleanedConfig, line...)
	}

	switch {
	case !foundStart && !foundEnd:
		return original, nil
	case !foundStart && foundEnd:
		return nil, errors.New("found end of injected background config but not start")
	case foundStart && !foundEnd:
		return nil, errors.New("found start of injected background config but not end")
	}

	return cleanedConfig, nil
}

const (
	bgConfigStartDelimiter = "### TOPBG START INJECTED CONFIG ###"
	bgConfigEndDelimiter   = "### TOPBG END INJECTED CONFIG ###"
	bgConfigTemplate       = `
{{ .ConfigStartPrefix }}
output * bg {{ .AbsoluteImagePath }} stretch
{{ .ConfigEndSuffix }}
`
)
