package logger

import (
	"log/slog"
	"os"
)

var instance *slog.Logger

func InitDefaultLogger() {
	instance = slog.New(slog.NewTextHandler(os.Stdout, nil))
}

// GetLogger returns the singleton logger instance.
//
// If instance is nil, it initializes the default logger.
func GetLogger() *slog.Logger {
	if instance == nil {
		InitDefaultLogger()
	}

	return instance
}
