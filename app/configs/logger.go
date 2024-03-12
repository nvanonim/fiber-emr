package configs

import (
	"github.com/nvanonim/fiber-emr/pkg/logger"
)

var log *logger.Logger

func SetupLogger() {
	log = logger.New()
}

// GetLogger returns the logger
func GetLogger() *logger.Logger {
	return log
}
