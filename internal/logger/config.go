package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Development bool   `koanf:"development"`
	Encoding    string `koanf:"encoding"`
	Level       string `koanf:"level"`
}

// getEncoder returns encoding level and encoder type.
func (cfg *Config) getEncoder() zapcore.Encoder {
	var encoderConfig zapcore.EncoderConfig
	if cfg.Development {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if cfg.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
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

	if !cfg.Development {
		options = append(options, zap.AddCaller())
		options = append(options, zap.AddStacktrace(zap.ErrorLevel))
	}

	return options
}
