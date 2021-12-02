package abstract

import (
	apmodelsv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

// AccountsRepo db access
type AccountsRepo interface {
	// Add add a new item
	Add(d *apmodelsv1.Account) (string, error)
	// Find find a account by id
	Find(id string) (*apmodelsv1.Account, error)
	// FindByUsername find a account by username
	FindByUsername(username string) (*apmodelsv1.Account, error)
	// FindByEmail find a account by email
	FindByEmail(email string) (*apmodelsv1.Account, error)
}
