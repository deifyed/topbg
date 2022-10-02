package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/deifyed/topbg/pkg/config"
	"github.com/google/uuid"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store the most recent topbg set background permanently",
	RunE: func(cmd *cobra.Command, args []string) error {
		fs := &afero.Afero{Fs: afero.NewOsFs()}

		temporaryImagePath := viper.GetString(config.TemporaryImageDir)
		imagesDirectory := viper.GetString(config.PermanentImageDir)

		img, err := findCurrentImageInDirectory(fs, path.Dir(temporaryImagePath), path.Base(temporaryImagePath))
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
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type image struct {
	Image     io.Reader
	Extension string
}

func findCurrentImageInDirectory(fs *afero.Afero, dir string, filename string) (image, error) {
	files, err := fs.ReadDir(dir)
	if err != nil {
		return image{}, fmt.Errorf("reading directory: %w", err)
	}

	var relevantFile os.FileInfo

	for _, file := range files {
		if extractFilename(file.Name()) != filename {
			continue
		}

		if relevantFile != nil && relevantFile.ModTime().Before(relevantFile.ModTime()) {
			continue
		}

		relevantFile = file

	}

	if relevantFile == nil {
		return image{}, errors.New("not found")
	}

	content, err := fs.ReadFile(path.Join(dir, relevantFile.Name()))
	if err != nil {
		return image{}, fmt.Errorf("reading file %s: %w", relevantFile.Name(), err)
	}

	return image{
		Image:     bytes.NewReader(content),
		Extension: strings.ReplaceAll(path.Ext(relevantFile.Name()), ".", ""),
	}, nil
}

func extractFilename(fullName string) string {
	parts := strings.Split(fullName, ".")

	return parts[0]
}
