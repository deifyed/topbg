package logging

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func ConfigureLogger(log *logrus.Logger, level string) {
	log.Formatter = &logrus.JSONFormatter{PrettyPrint: true}
	log.Out = os.Stdout

	switch strings.ToLower(level) {
	case "debug":
		log.Level = logrus.DebugLevel
	case "info":
		log.Level = logrus.InfoLevel
	default:
		log.Level = logrus.InfoLevel
	}
}
