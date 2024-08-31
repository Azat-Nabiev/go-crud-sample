package config

import (
	"go.uber.org/zap"
	"log"
)

func InitLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Exception during logger initizalization")
	}
	return logger.Sugar()
}
