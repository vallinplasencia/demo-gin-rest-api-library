package mysql

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	User     string `envconfig:"DB_MYSQL_USER"`
	Password string `envconfig:"DB_MYSQL_PASSWORD"`
	// Address formato "127.0.0.1:3306"
	Address string `envconfig:"DB_MYSQL_ADDRESS"`
	DBName  string `envconfig:"DB_MYSQL_DBNAME"`
}

// ConfigFromEnv ...
func ConfigFromEnv(prefix string) (*config, error) {
	conf := config{}
	e := envconfig.Process(prefix, &conf)
	return &conf, e
}
