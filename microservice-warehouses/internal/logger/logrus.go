package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/config"
	log "github.com/sirupsen/logrus"
)

var file *os.File // saved file handle for closing later

// SetupLogger creates and configures a new Logrus logger based on the provided configuration.
func SetupLogger(cfg *config.LoggerConfig) error {

	setLogrusLevel(cfg.Level)

	setFormatter(cfg.Format)

	err := setOutput(cfg.Output)
	if err != nil {
		return fmt.Errorf("failed to set log output: %w", err)
	}

	return nil
}

func setLogrusLevel(level int) {
	switch level {
	case config.LevelDebug:
		log.SetLevel(log.DebugLevel)
	case config.LevelInfo:
		log.SetLevel(log.InfoLevel)
	case config.LevelWarn:
		log.SetLevel(log.WarnLevel)
	case config.LevelError:
		log.SetLevel(log.ErrorLevel)
	case config.LevelFatal:
		log.SetLevel(log.FatalLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

func setFormatter(format string) {
	switch format {
	case config.TextFormat:
		log.SetFormatter(&log.TextFormatter{})
	case config.JSONFormat:
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}
}

func setOutput(output string) error {
	if output == config.OutputFile {
		file, err := createLogFile()
		if err != nil {
			log.Errorf("Failed to create log file: %v", err)
			log.SetOutput(os.Stdout)
			return fmt.Errorf("failed to create log file: %w", err)
		}
		log.SetOutput(file)
		return nil
	}

	log.SetOutput(os.Stdout)

	return nil
}

func createLogFile() (*os.File, error) {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		if err := os.Mkdir("logs", 0755); err != nil {
			return nil, fmt.Errorf("failed to create logs directory: %w", err)
		}
	}

	fileName := fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02_15-04-05"))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}
	return file, nil
}

func Close() error {
	if file != nil {
		return file.Close()
	}
	return nil
}
