package config

import (
	"github.com/sadqx/aenema/internal/logger"
	"github.com/sadqx/aenema/internal/tracer"
)

// Default return default configuration.
func Default() Config {
	return Config{
		Logger: &logger.Config{
			Development: true,
			Encoding:    "console",
			Level:       "info",
		},
		Tracer: &tracer.Config{
			Host:       "",
			Port:       6831,
			SampleRate: 0.1,
			IsEnabled:  false,
		},
	}
}
