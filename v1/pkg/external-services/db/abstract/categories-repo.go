package abstract

import (
	apmodelsv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

// CategoriesRepo db access
type CategoriesRepo interface {
	// Add add a new item
	Add(d *apmodelsv1.Category) (string, error)
	// Find return a item by id.
	//
	// If id not found return ErrorNoItems
	Find(id string) (*apmodelsv1.Category, error)
}
