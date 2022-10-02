package cmd

import (
	"errors"

	"github.com/deifyed/topbg/cmd/set"
	"github.com/deifyed/topbg/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SetCmd represents the Set command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set background to a random image",
	Long:  `Grabs a random image from the configured list of subreddits`,
	Args:  cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(viper.GetStringSlice(config.Subreddits)) == 0 {
			return errors.New("no subreddits configured")
		}

		return nil
	},
	RunE: set.RunE(log, fs),
}

func init() {
	rootCmd.AddCommand(SetCmd)

	viper.SetDefault(
		config.Subreddits,
		[]string{"earthporn", "abandonedporn", "dalle2", "midjourney"},
	)

	SetCmd.Flags().StringArrayVarP(
		&setCmdOpts.Subreddits,
		config.Subreddits,
		"s",
		viper.GetStringSlice(config.Subreddits),
		"Subreddits to gather images from",
	)

	viper.BindPFlag(config.Subreddits, SetCmd.Flags().Lookup(config.Subreddits))
}

var setCmdOpts = struct {
	Subreddits []string
}{}
