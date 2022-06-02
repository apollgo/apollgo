//nolint:gomnd
package config

import (
	"github.com/sadqx/aenema/internal/logger"
	"github.com/sadqx/aenema/internal/tracer"
)

func Default() Config {
	return Config{
		Logger: &logger.Config{
			Encode: "console",
			Level:  "info",
			Debug:  true,
		},
		Tracer: &tracer.Config{
			Host:       "",
			Port:       6831,
			IsEnabled:  false,
			SampleRate: 0.1,
		},
	}
}
