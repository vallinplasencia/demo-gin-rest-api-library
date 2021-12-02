package v1

import (
	"github.com/kelseyhightower/envconfig"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	aputil "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/util"
)

// Handlers ...
type Handlers struct {
	Books    *BookHandler
	Accounts *AccountsHandler
}

func New(c *config, db apdbabstract.DB) *Handlers {
	base := &Base{
		DB:  db,
		env: c.Env,
	}
	books := &BookHandler{
		Base: base,
	}
	accounts := &AccountsHandler{
		Base: base,
	}
	return &Handlers{
		Books:    books,
		Accounts: accounts,
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
	case string(aputil.EnvProduction):
		conf.Env = aputil.EnvProduction
	case string(aputil.EnvDevelop):
		conf.Env = aputil.EnvDevelop
	}
	return &conf, e
}
