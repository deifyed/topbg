package cmd

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func createLogger() *logrus.Logger {
	log := logrus.Logger{}

	log.Formatter = &logrus.JSONFormatter{PrettyPrint: true}
	log.Out = os.Stdout
	log.Level = logrus.InfoLevel

	switch strings.ToLower(os.Getenv("TOPBG_LOG_LEVEL")) {
	case "debug":
		log.Level = logrus.DebugLevel
	default:
		break
	}

	return &log
}
