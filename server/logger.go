package server

import (
	"log"
	"os"

	"go.uber.org/zap"
)

func newLogger() *zap.Logger {
	var loggerFactory func(...zap.Option) (*zap.Logger, error)

	if appEnv := os.Getenv("APP_ENV"); appEnv == "production" {
		loggerFactory = zap.NewProduction
	} else {
		loggerFactory = zap.NewDevelopment
	}

	logger, err := loggerFactory()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	return logger
}
