package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

type SharedConfig struct {
	// LogLevel is the log level to use
	LogLevel string `split_words:"true" default:"info"`
}

func initLoggingConfig(config *SharedConfig) error {
	logLevelMap := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}

	level, levelValid := logLevelMap[config.LogLevel]

	if !levelValid {
		return fmt.Errorf("invalid log level: %s", config.LogLevel)
	}

	stderr := os.Stderr
	writer := colorable.NewColorable(stderr)

	slog.SetDefault(
		slog.New(
			tint.NewHandler(
				writer,
				&tint.Options{
					Level:   level,
					NoColor: !isatty.IsTerminal(stderr.Fd()),
				},
			),
		),
	)

	return nil
}

func loadConfig[T any]() (*T, error) {
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		slog.Error("Error loading .env file", "err", err)
	}

	var config T

	err = envconfig.Process("monticola", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
