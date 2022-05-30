package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(cfg *Config) *zap.Logger {
	return zap.New(
		zapcore.NewCore(
			cfg.getEncoder(),
			cfg.getWriteSyncer(),
			cfg.getLoggerLevel(),
		),
		cfg.getOptions()...,
	)
}
