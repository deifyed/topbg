package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
)

func createLogger() *logrus.Logger {
	log := logrus.Logger{}

	log.Formatter = &logrus.JSONFormatter{PrettyPrint: true}
	log.Out = os.Stdout
	log.Level = logrus.InfoLevel

	switch os.Getenv("TOPBG_DEBUG_LEVEL") {
	case "DEBUG":
		log.Level = logrus.DebugLevel
	default:
		break
	}

	return &log
}
