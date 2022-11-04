package wm

import (
	"io"
	"strings"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestInjectBackgroundConfig(t *testing.T) {
	t.Skip("Not relevant until functionality is implemented")

	testCases := []struct {
		name          string
		withConfig    io.Reader
		withImagePath string
	}{
		{
			name:          "Should work when adding to config without previously set background",
			withConfig:    strings.NewReader(configSnippet),
			withImagePath: "/home/user/images/mock-img.jpg",
		},
		{
			name:          "Should work when adding to config with previously set background",
			withConfig:    strings.NewReader(configSnippetWithBackground),
			withImagePath: "/home/user/images/mock-img.jpg",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			fs := &afero.Afero{Fs: afero.NewMemMapFs()}

			configPath := "/home/user/.config/sway/config"

			err := fs.WriteReader(configPath, tc.withConfig)
			assert.NoError(t, err)

			err = injectBackgroundConfig(fs, configPath, tc.withImagePath)
			assert.NoError(t, err)

			g := goldie.New(t)

			g.Assert(t, t.Name(), mustRead(t, fs, configPath))
		})
	}
}

func mustRead(t *testing.T, fs *afero.Afero, path string) []byte {
	t.Helper()

	data, err := fs.ReadFile(path)
	assert.NoError(t, err)

	return data
}

const (
	configSnippet = `#
# Status Bar:
#
# Read man 5 sway-bar for more information about this section.
bar {
    position top

    # When the status_command prints a new line to stdout, swaybar updates.
    # The default just shows the current date and time.
    #status_command while ~/.config/sway/status.sh; do sleep 1; done
    #status_command while ~/.config/sway/status_bar.py; do sleep 1; done
    status_command while status; do sleep 1; done

    colors {
			# Status line font
			statusline #a89984
			# Status line background
			background #282828

			# border main font
			focused_workspace #ebdbb2 #ebdbb2 #282828
			inactive_workspace #928374 #928374 #282828
			urgent_workspace #fb4934 #928374 #fb4934
    }
}

include /etc/sway/config.d/*
include /home/user/.config/sway/config.d/*
`
	configSnippetWithBackground = configSnippet + `

### TOPBG START INJECTED CONFIG ###
output * bg /home/user/images/old-mock-img.jpg stretch
### TOPBG END INJECTED CONFIG ###
`
)
