package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/config"
	"github.com/rs/zerolog"
)

var file *os.File // saved file handle for closing later

// SetupLogger creates and configures a new apex/log logger based on the provided configuration.
func SetupLogger(cfg *config.LoggerConfig) (*zerolog.Logger, error) {
	level := getLogLevel(cfg.Level)

	output, err := getLogOutput(cfg.Output)
	if err != nil {
		return nil, fmt.Errorf("failed to set log output: %w", err)
	}

	formatter := getFormatter(cfg.Format, output)

	log := zerolog.New(formatter).Level(level).With().Timestamp().Logger()
	return &log, nil
}

// logLevels maps custom log levels to zerolog.Level.
var logLevels = map[int]zerolog.Level{
	config.LevelDebug: zerolog.DebugLevel,
	config.LevelInfo:  zerolog.InfoLevel,
	config.LevelWarn:  zerolog.WarnLevel,
	config.LevelError: zerolog.ErrorLevel,
	config.LevelFatal: zerolog.FatalLevel,
}

// getLogLevel maps the custom log level to zerolog.Level.
func getLogLevel(level int) zerolog.Level {
	if logLevel, exists := logLevels[level]; exists {
		return logLevel
	}
	return zerolog.InfoLevel // default to InfoLevel if not found
}

// getFormatter returns the appropriate log formatter based on the specified format.
func getFormatter(format string, output io.Writer) io.Writer {
	switch format {
	case config.TextFormat:
		return zerolog.ConsoleWriter{
			Out: output,
		}
	case config.JSONFormat:
		return output
	default:
		return output
	}
}

// getLogOutput sets the log output based on the provided output string.
func getLogOutput(output string) (io.Writer, error) {
	if output == config.OutputFile {
		file, err := createLogFile()
		if err != nil {
			log.Printf("Failed to create log file: %v\n", err)
			return nil, fmt.Errorf("failed to create log file: %w", err)
		}
		return file, nil
	}

	return os.Stdout, nil
}

// createLogFile creates a new log file with a timestamped name in the "logs" directory.
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

// Close closes the log file if it was opened.
func Close() error {
	if file != nil {
		return file.Close()
	}
	return nil
}
