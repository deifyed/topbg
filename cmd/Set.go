package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/deifyed/topbg/pkg/reddit"
	"github.com/deifyed/topbg/pkg/wm"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// SetCmd represents the Set command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set background to a random image",
	Long:  `Grabs a random image from the configured list of subreddits`,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		subreddits := []string{"earthporn", "abandonedporn", "dalle2", "midjourney"}
		fs := &afero.Afero{Fs: afero.NewOsFs()}
		log := createLogger()

		imageURLs := make([]string, 0)

		for _, subreddit := range subreddits {
			urls, err := reddit.GetSubreddit(log, subreddit)
			if err != nil {
				return fmt.Errorf("fetching subreddit %s: %w", subreddit, err)
			}

			imageURLs = append(imageURLs, urls...)
		}

		log.Debugf("Found URLs: %v", imageURLs)

		rand.Seed(time.Now().Unix())
		relevantURL := imageURLs[rand.Intn(len(imageURLs))]

		log.Debugf("Chose URL %s", relevantURL)

		image, err := reddit.DownloadImage(relevantURL)
		if err != nil {
			return fmt.Errorf("downloading image: %w", err)
		}

		err = wm.SetBackground(fs, image.Type, image.Image)
		if err != nil {
			return fmt.Errorf("setting background: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(SetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// SetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// SetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
