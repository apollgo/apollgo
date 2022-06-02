package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Encode string `koanf:"encode"`
	Level  string `koanf:"level"`
	Debug  bool   `koanf:"debug"`
}

// getEncoder returns encoding level and encoder type.
func (cfg *Config) getEncoder() zapcore.Encoder {
	var encoderConfig zapcore.EncoderConfig
	if cfg.Debug {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	switch strings.ToLower(cfg.Encode) {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	return encoder
}

// getWriteSyncer warps writer with a mutex.
func (cfg *Config) getWriteSyncer() zapcore.WriteSyncer {
	return zapcore.Lock(os.Stdout)
}

// getLoggerLevel parses logger lever and returns it.
func (cfg *Config) getLoggerLevel() zap.AtomicLevel {
	var level zapcore.Level

	if err := level.Set(cfg.Level); err != nil {
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	return zap.NewAtomicLevelAt(level)
}

// getOptions return extra option for stacktrace on non-development env.
func (cfg *Config) getOptions() []zap.Option {
	options := make([]zap.Option, 0)

	if !cfg.Debug {
		options = append(options, zap.AddCaller())
		options = append(options, zap.AddStacktrace(zap.ErrorLevel))
	}

	return options
}
