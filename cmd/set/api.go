package set

import (
	"bytes"
	"fmt"
	"math/rand"
	"path"
	"strings"
	"time"

	"github.com/deifyed/topbg/pkg/config"
	"github.com/deifyed/topbg/pkg/images"
	"github.com/deifyed/topbg/pkg/logging"
	"github.com/deifyed/topbg/pkg/reddit"
	"github.com/deifyed/topbg/pkg/wm"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE(log logging.Logger, fs *afero.Afero, targetIndex *int) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if *targetIndex < 0 {
			return setRandomBackground(log, fs)
		}

		err := setBackgroundToImageByIndex(fs, *targetIndex)
		if err != nil {
			return fmt.Errorf("setting background to image index: %w", err)
		}

		return nil
	}
}

func setRandomBackground(log logging.Logger, fs *afero.Afero) error {
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

func setBackgroundToImageByIndex(fs *afero.Afero, index int) error {
	paths, err := images.ListPaths(fs)
	if err != nil {
		return fmt.Errorf("listing paths: %w", err)
	}

	relevantPath := paths[index]

	content, err := fs.ReadFile(relevantPath)
	if err != nil {
		return fmt.Errorf("reading image %s: %w", relevantPath, err)
	}

	err = wm.SetBackground(
		fs,
		strings.ReplaceAll(path.Ext(relevantPath), ".", ""),
		bytes.NewReader(content),
	)
	if err != nil {
		return fmt.Errorf("setting background to %d at %s: %w", index, relevantPath, err)
	}

	return nil
}
