package local

import (
	"github.com/kelseyhightower/envconfig"
)

// Config ...
type Config struct {
	// DestinationTarget directorio a subir el archivo
	DestinationTarget string `envconfig:"DESTINATION_TARGET"`
}

// ConfigFromEnv ...
func ConfigFromEnv(prefix string) (*Config, error) {
	conf := Config{}
	e := envconfig.Process(prefix, &conf)
	return &conf, e
}
