package logging

import (
	"os"
	"strings"

	"github.com/deifyed/topbg/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger() *logrus.Logger {
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
		log.Level = logrus.InfoLevel
	}

	return &log
}
