package store

import (
	"fmt"
	"path"

	"github.com/deifyed/topbg/pkg/config"
	"github.com/deifyed/topbg/pkg/logging"
	"github.com/deifyed/topbg/pkg/wm"
	"github.com/google/uuid"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE(log logging.Logger, fs *afero.Afero) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		temporaryImagePath := viper.GetString(config.TemporaryImageDir)
		imagesDirectory := viper.GetString(config.PermanentImageDir)

		img, err := findCurrentImageInDirectory(fs, temporaryImagePath, wm.TemporaryFilename)
		if err != nil {
			return fmt.Errorf("acquiring current image: %w", err)
		}

		err = fs.WriteReader(
			path.Join(imagesDirectory, fmt.Sprintf("%s.%s", uuid.New().String(), img.Extension)),
			img.Image,
		)
		if err != nil {
			return fmt.Errorf("writing image: %w", err)
		}

		return nil
	}
}
