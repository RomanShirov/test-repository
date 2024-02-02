package main

import (
	"test-repository/logger"
)

func main() {
	l := logger.New()

	l.Log.Info("Hello world")
}
