package s3

import (
	"github.com/kelseyhightower/envconfig"
)

// config ...
type config struct {
	// DestinationTarget bucket s3 a subir el archivo
	DestinationTarget string `envconfig:"DESTINATION_TARGET"`
}

// ConfigFromEnv ...
func ConfigFromEnv(prefix string) (*config, error) {
	conf := config{}
	e := envconfig.Process(prefix, &conf)
	return &conf, e
}
