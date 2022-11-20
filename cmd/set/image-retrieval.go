package set

import (
	"fmt"
	"math/rand"
	"path"
	"time"

	"github.com/deifyed/topbg/pkg/config"
	"github.com/deifyed/topbg/pkg/images"
	"github.com/deifyed/topbg/pkg/reddit"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

func getRandomImagePath(log logger, fs *afero.Afero) (string, error) {
	img, err := getRandomImage(log)
	if err != nil {
		return "", fmt.Errorf("getting random image: %w", err)
	}

	targetPath := path.Join(viper.GetString(config.TemporaryImageDir), fmt.Sprintf("topbg-temp-image.%s", img.Type))

	err = fs.WriteReader(targetPath, img.Image)
	if err != nil {
		return "", fmt.Errorf("writing image to %s: %w", targetPath, err)
	}

	return targetPath, nil
}

func getRandomImage(log logger) (reddit.Image, error) {
	subreddits := viper.GetStringSlice(config.Subreddits)
	rand.Seed(time.Now().Unix())
	relevantSubreddit := subreddits[rand.Intn(len(subreddits))]

	log.Debugf("Picking from following subreddits: %v", subreddits)
	log.Debugf("Picked subreddit: %s", relevantSubreddit)

	urls, err := reddit.GetSubreddit(log, relevantSubreddit)

	log.Debugf("Found URLs: %v", urls)

	relevantURL := urls[rand.Intn(len(urls))]

	log.Debugf("Chose URL %s", relevantURL)

	image, err := reddit.DownloadImage(relevantURL)
	if err != nil {
		return reddit.Image{}, fmt.Errorf("downloading image: %w", err)
	}

	return image, nil
}

func getImagePathByIndex(fs *afero.Afero, index int) string {
	imagePaths, err := images.ListPaths(fs)
	if err != nil {
		return ""
	}

	return imagePaths[index]
}
