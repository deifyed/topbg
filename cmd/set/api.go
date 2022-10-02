package set

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/deifyed/topbg/pkg/config"
	"github.com/deifyed/topbg/pkg/logging"
	"github.com/deifyed/topbg/pkg/reddit"
	"github.com/deifyed/topbg/pkg/wm"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE(log logging.Logger, fs *afero.Afero) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		subreddits := viper.GetStringSlice(config.Subreddits)
		imageURLs := make([]string, 0)

		log.Debugf("Picking from following subreddits: %v", subreddits)

		for _, subreddit := range subreddits {
			log.Debugf("Fetching %s URLs", subreddit)

			urls, err := reddit.GetSubreddit(subreddit)
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
	}
}
