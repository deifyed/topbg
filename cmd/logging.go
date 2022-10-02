package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/deifyed/topbg/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func createLogger() *logrus.Logger {
	log := logrus.Logger{}

	log.Formatter = &logrus.JSONFormatter{PrettyPrint: true}
	log.Out = os.Stdout
	log.Level = logrus.InfoLevel

	targetLevel := viper.GetString(config.LogLevel)

	switch strings.ToLower(targetLevel) {
	case "debug":
		log.Level = logrus.DebugLevel
	case "info":
		log.Level = logrus.InfoLevel
	default:
		panic(fmt.Sprintf("Unknown log level %s", targetLevel))
	}

	return &log
}
