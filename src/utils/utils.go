package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

// GetLogger configures the logging and return a logger
func GetLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	return logger
}

