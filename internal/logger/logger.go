package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger builds a logger instance with specific configuration.
func NewLogger(cfg *Config) *zap.Logger {
	return zap.New(
		zapcore.NewCore(
			cfg.getEncoder(),
			cfg.getWriteSyncer(),
			cfg.getLoggerLevel(),
		),
		cfg.getOptions()...,
	)
}
