package v1

import (
	"github.com/kelseyhightower/envconfig"

	apauthtokenabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/auth/access-token/abstract"
	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	apstoreabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/store/abstract"
	aputil "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/util"
)

// Handlers ...
type Handlers struct {
	Books    *BookHandler
	Accounts *AccountsHandler
}

func New(c *config, db apdbabstract.DB, authToken apauthtokenabstract.Token, store apstoreabstract.Store) *Handlers {
	base := &base{
		db:         db,
		env:        c.Env,
		token:      authToken,
		storeFiles: store,
	}
	books := &BookHandler{
		base: base,
	}
	accounts := &AccountsHandler{
		base: base,
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
