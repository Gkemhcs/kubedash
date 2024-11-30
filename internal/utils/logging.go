package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	*logrus.Logger
}

func (logger *LoggerConfig) InitLogger() *LoggerConfig {
	logger.Logger = logrus.New() // This in
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel)

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Warn("Failed to log to file, using default stderr")
	}
	return logger
}
