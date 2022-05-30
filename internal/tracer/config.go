package tracer

type Config struct {
	SampleRate float64 `koanf:"sample-rate"`
	IsEnabled  bool    `koanf:"enabled"`
	Host       string  `koanf:"host"`
	Port       int     `koanf:"port"`
}
