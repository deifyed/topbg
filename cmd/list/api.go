package list

import (
	"fmt"
	"sort"

	"github.com/deifyed/topbg/pkg/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE(fs *afero.Afero) func(*cobra.Command, []string) error {
	return func(c *cobra.Command, s []string) error {
		files, err := fs.ReadDir(viper.GetString(config.PermanentImageDir))
		if err != nil {
			return fmt.Errorf("listing files in directory: %w", err)
		}

		imageNames := make([]string, len(files))

		for index, file := range files {
			imageNames[index] = file.Name()
		}

		imageNames = sort.StringSlice(imageNames)

		for index, item := range imageNames {
			fmt.Printf("[%d] %s\n", index, item)
		}

		return nil
	}
}
