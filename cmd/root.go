package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/deifyed/topbg/cmd/set"
	"github.com/deifyed/topbg/pkg/config"
	"github.com/deifyed/topbg/pkg/logging"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	log = logging.NewLogger()
	fs  = &afero.Afero{Fs: afero.NewOsFs()}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "topbg",
	Short:        "Grabs top images from subreddits and sets the wallpaper with them",
	Long:         `TopBG grabs a random image from the top posts of configured subreddits and sets it as the desktop wallpaper`,
	SilenceUsage: true,
	RunE:         set.RunE(log, fs, intPtr(-1)),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default $HOME/.topbg.yaml)")

	rootCmd.PersistentFlags().StringVarP(&rootCmdOpts.LogLevel, "log-level", "l", defaultLogLevel, "Set log level")
	viper.BindEnv(config.LogLevel)
	viper.BindPFlag(config.LogLevel, rootCmd.Flags().Lookup("log-level"))
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("acquiring home directory: %s", err.Error()))
	}

	cfgDir := path.Join(home, ".config", "topbg")

	viper.AddConfigPath(home)
	viper.AddConfigPath(cfgDir)
	viper.SetConfigName("topbg")
	viper.SetEnvPrefix("topbg")

	// Defaults
	viper.SetDefault(config.TemporaryImageDir, os.TempDir())
	viper.SetDefault(config.PermanentImageDir, path.Join(cfgDir, "images"))

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Sprintf("reading configuration: %s", err.Error()))
		}
	}
}

var configFile string
var rootCmdOpts struct {
	LogLevel string
}

const defaultLogLevel = "info"

func intPtr(i int) *int {
	return &i
}
