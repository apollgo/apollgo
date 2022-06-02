//nolint:gomnd
package config

import (
	"github.com/apollgo/apollgo/internal/logger"
	"github.com/apollgo/apollgo/internal/tracer"
)

// Default returns dummy configuration.
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
