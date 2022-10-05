package list

import (
	"fmt"
	"path"

	"github.com/deifyed/topbg/pkg/images"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func RunE(fs *afero.Afero) func(*cobra.Command, []string) error {
	return func(c *cobra.Command, s []string) error {
		paths, err := images.ListPaths(fs)
		if err != nil {
			return fmt.Errorf("listing paths: %w", err)
		}

		for index, p := range paths {
			fmt.Printf("[%d] %s\n", index, path.Base(p))
		}

		return nil
	}
}
