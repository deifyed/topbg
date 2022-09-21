package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
)

func createLogger() *logrus.Logger {
	log := logrus.Logger{}

	log.Formatter = &logrus.JSONFormatter{PrettyPrint: true}
	log.Out = os.Stdout

	log.Level = logrus.DebugLevel

	return &log
}
