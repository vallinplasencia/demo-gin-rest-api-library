package dgjwt

import (
	"github.com/kelseyhightower/envconfig"
)

// Config ...
type config struct {
	AccessTokenSecretKey string `envconfig:"ACCESS_TOKEN_SECRET_KEY"`
	AccessTokenAudience  string `envconfig:"ACCESS_TOKEN_AUDIENCE"`
	AccessTokenIssuer    string `envconfig:"ACCESS_TOKEN_ISSUER"`
	AccessTokenLive      int64  `envconfig:"ACCESS_TOKEN_LIVE"`

	RefreshTokenSecretKey string `envconfig:"REFRESH_TOKEN_SECRET_KEY"`
	RefreshTokenLive      int64  `envconfig:"REFRESH_TOKEN_LIVE"`

	// URLServerStoreMedias url del servidor donde se alojan los archivos subidos
	URLServerStoreMedias string `envconfig:"URL_SERVER_STORE_MEDIAS"`
}

// ConfigFromEnv ...
func ConfigFromEnv(prefix string) (*config, error) {
	conf := config{}
	e := envconfig.Process(prefix, &conf)
	return &conf, e
}
