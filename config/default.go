//nolint:gomnd
package config

import (
	"github.com/apollgo/apollgo/logger"
	"github.com/apollgo/apollgo/tracer"
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
