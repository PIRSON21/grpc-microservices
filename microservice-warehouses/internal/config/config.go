package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Environment constants
const (
	EnvDevelopment = "dev"  // Local development
	EnvTesting     = "test" // Testing environment
	EnvProduction  = "prod" // Production environment
)

// Config holds the application configuration.
type Config struct {
	Addr         string       `mapstructure:"addr"`
	Env          string       `mapstructure:"env"`
	LoggerConfig LoggerConfig `mapstructure:",squash"`
}

func (c *Config) validate() error {
	if c.Env != EnvDevelopment && c.Env != EnvTesting && c.Env != EnvProduction {
		return fmt.Errorf("invalid env: %s", c.Env)
	}

	if err := c.LoggerConfig.validate(); err != nil {
		return fmt.Errorf("invalid logger config: %w", err)
	}

	return nil
}

// Log level constants
const (
	LevelDebug = 1 // Debug level
	LevelInfo  = 2 // Info level
	LevelWarn  = 3 // Warn level
	LevelError = 4 // Error level
	LevelFatal = 5 // Fatal level
)

// Log format constants
const (
	TextFormat = "text" // Plain text format
	JSONFormat = "json" // JSON format
)

// Output constants
const (
	OutputConsole = "console" // Output to console
	OutputFile    = "file"    // Output to file
)

// LoggerConfig holds the logger configuration.
type LoggerConfig struct {
	Level  int    `mapstructure:"level"`  // Log level
	Format string `mapstructure:"format"` // Log format
	Output string `mapstructure:"output"` // Log output
}

func (c *LoggerConfig) validate() error {
	if c.Level < LevelDebug || c.Level > LevelFatal {
		return fmt.Errorf("invalid log level: %d", c.Level)
	}
	return nil
}

// validate checks if the configuration values are valid.

// MustLoadConfig loads the configuration from the specified path or environment variables.
func MustLoadConfig(path *string) *Config {
	cfg, err := loadConfig(path)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	if err := cfg.validate(); err != nil {
		panic(fmt.Errorf("invalid config: %w", err))
	}

	return cfg
}

// loadConfig loads the configuration from a file or environment variables.
func loadConfig(path *string) (*Config, error) {
	loadDefaultValuesForViper()

	if *path != "" {
		cfg, err := loadConfigFromFile(*path)
		if err != nil {
			return nil, fmt.Errorf("failed to load config from file: %w", err)
		}

		return cfg, nil
	}

	cfg, err := loadConfigFromEnv()
	if err != nil {
		return nil, fmt.Errorf("failed to load config from environment variables: %w", err)
	}

	return cfg, nil
}

// loadConfigFromFile loads the configuration from a file.
func loadConfigFromFile(path string) (*Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// loadConfigFromEnv loads the configuration from environment variables.
func loadConfigFromEnv() (*Config, error) {
	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// loadDefaultValuesForViper sets default values for configuration keys in Viper.
func loadDefaultValuesForViper() {
	viper.SetDefault("addr", ":8080")
	viper.SetDefault("env", EnvProduction)
	viper.SetDefault("level", LevelInfo)
	viper.SetDefault("format", TextFormat)
	viper.SetDefault("output", OutputConsole)
}
