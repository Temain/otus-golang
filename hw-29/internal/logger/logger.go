package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(logFile string, level string) *logrus.Logger {
	var logger = logrus.New()
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, f)
	logrus.SetOutput(mw)

	var loggerLevel logrus.Level
	switch level {
	case "info":
		loggerLevel = logrus.InfoLevel
		break
	case "debug":
		loggerLevel = logrus.DebugLevel
		break
	case "warn":
		loggerLevel = logrus.WarnLevel
		break
	case "error":
		loggerLevel = logrus.ErrorLevel
		break
	}
	logger.SetLevel(loggerLevel)
	return logger
}
