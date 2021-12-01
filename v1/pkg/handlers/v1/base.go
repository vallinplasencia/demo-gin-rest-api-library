package v1

import (
	"github.com/kelseyhightower/envconfig"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	aputil "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/util"
)

// Base ...
type Base struct {
	env aputil.EnvType
	DB  apdbabstract.DB
}

// Handlers ...
type Handlers struct {
	Book *BookHandler
}

func New(c *config, db apdbabstract.DB) *Handlers {
	base := &Base{
		DB:  db,
		env: c.Env,
	}
	book := &BookHandler{
		Base: base,
	}
	return &Handlers{
		Book: book,
	}
}

// ======= config ======= //

type config struct {
	Env aputil.EnvType
}

// ConfigFromEnv ...
func ConfigFromEnv(prefix string) (*config, error) {
	temp := struct {
		Env string `envconfig:"APP_MODE"`
	}{}
	e := envconfig.Process(prefix, &temp)

	conf := config{
		Env: aputil.EnvProduction,
	}
	switch temp.Env {
	case "production":
		conf.Env = aputil.EnvProduction
	case "develop":
		conf.Env = aputil.EnvDevelop
	}
	return &conf, e
}
