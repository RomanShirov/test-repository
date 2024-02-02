package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Log *logrus.Logger
}

func New() *Logger {
	logger := logrus.New()

	return &Logger{
		Log: logger,
	}
}
