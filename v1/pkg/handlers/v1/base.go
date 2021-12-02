package v1

import (
	"errors"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	apv1models "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
	aputil "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/util"
)

var ErrorRespIsEmpty error = errors.New("response from db is nil")

const (
	concFindAccountByEmail    aputil.ConcName = "find-account-by-email"
	concFindAccountByUsername aputil.ConcName = "find-account-by-username"
)

// Base ...
type Base struct {
	env aputil.EnvType
	DB  apdbabstract.DB
}

// findAccountByEmail recupera una cuenta por su email
func (b *Base) findAccountByEmail(email string, ch chan *aputil.ConcurrencyData, out *apv1models.Account) {
	d, e := b.DB.Accounts().FindByEmail(email)
	cd := &aputil.ConcurrencyData{
		Err:  nil,
		Name: concFindAccountByEmail,
	}
	switch {
	case e != nil:
		cd.Err = e
	case d == nil:
		cd.Err = ErrorRespIsEmpty
	default:
		*out = *d
	}
	ch <- cd
}

// findAccountByUsername recupera una cuenta por su username
func (b *Base) findAccountByUsername(username string, ch chan *aputil.ConcurrencyData, out *apv1models.Account) {
	d, e := b.DB.Accounts().FindByUsername(username)
	cd := &aputil.ConcurrencyData{
		Err:  nil,
		Name: concFindAccountByUsername,
	}
	switch {
	case e != nil:
		cd.Err = e
	case d == nil:
		cd.Err = ErrorRespIsEmpty
	default:
		*out = *d
	}
	ch <- cd
}
