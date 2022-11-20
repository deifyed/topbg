package set

import (
	"fmt"

	"github.com/deifyed/topbg/pkg/config"
	"github.com/deifyed/topbg/pkg/wm"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE(log logger, fs *afero.Afero, opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var (
			img image
			err error
		)

		fmt.Printf("%+v\n", opts)

		if opts.Index < 0 {
			img.AbsolutePath, err = getRandomImagePath(log, fs)
			if err != nil {
				return fmt.Errorf("getting random image: %w", err)
			}
		} else {
			img.AbsolutePath = getImagePathByIndex(fs, opts.Index)
		}

		imgContent, err := img.Content(fs)
		if err != nil {
			return fmt.Errorf("getting image content: %w", err)
		}

		err = wm.SetBackground(fs, img.Type(), imgContent)
		if err != nil {
			return fmt.Errorf("setting background: %w", err)
		}

		if !opts.Permanent {
			return nil
		}

		if opts.Index < 0 {
			return fmt.Errorf("cannot set permanent background from random image. Store it first")
		}

		cfgPath, err := ensureAbs(viper.GetString(config.I3TopBGConfigurationPath))
		if err != nil {
			return fmt.Errorf("getting absolute path: %w", err)
		}

		log.Debugf("configuration path: %s", cfgPath)

		err = wm.InjectBackgroundConfig(fs, cfgPath, img.AbsolutePath)
		if err != nil {
			return fmt.Errorf("injecting background config: %w", err)
		}

		return nil
	}
}

func PreRunE(log logger, fs *afero.Afero, opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(opts.Subreddits) < 1 {
			return fmt.Errorf("no subreddits provided")
		}

		return nil
	}
}
