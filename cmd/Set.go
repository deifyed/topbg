package cmd

import (
	"github.com/deifyed/topbg/cmd/set"
	"github.com/deifyed/topbg/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setCmd represents the Set command
var (
	setCmdOpts = set.Options{Index: -1}
	setCmd     = &cobra.Command{
		Use:     "set",
		Short:   "Set background to a random image",
		Long:    `Grabs a random image from the configured list of subreddits`,
		Args:    cobra.ExactArgs(0),
		PreRunE: set.PreRunE(log, fs, &setCmdOpts),
		RunE:    set.RunE(log, fs, &setCmdOpts),
	}
)

func init() {
	rootCmd.AddCommand(setCmd)

	viper.SetDefault(
		config.Subreddits,
		[]string{"earthporn", "abandonedporn", "dalle2", "midjourney"},
	)

	setCmd.Flags().StringArrayVarP(
		&setCmdOpts.Subreddits,
		config.Subreddits,
		"s",
		viper.GetStringSlice(config.Subreddits),
		"Subreddits to gather images from",
	)
	err := viper.BindPFlag(config.Subreddits, setCmd.Flags().Lookup(config.Subreddits))
	cobra.CheckErr(err)

	setCmd.Flags().IntVarP(&setCmdOpts.Index, "index", "i", setCmdOpts.Index, "Set wallpaper by index to stored image")
	setCmd.Flags().BoolVarP(&setCmdOpts.Permanent, "permanent", "p", false, "Set wallpaper permanently")
}
