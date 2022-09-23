package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "topbg",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	SilenceUsage: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	viper.BindEnv("LogLevel")
	viper.BindPFlag("LogLevel", rootCmd.Flags().Lookup("log-level"))
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("acquiring home directory: %s", err.Error()))
	}

	viper.AddConfigPath(home)
	viper.AddConfigPath(path.Join(home, ".config", "topbg"))
	viper.SetConfigName("topbg")
	viper.SetEnvPrefix("topbg")

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
