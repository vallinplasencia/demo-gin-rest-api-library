package abstract

import (
	apmodelsv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

// BooksRepo db access
type BooksRepo interface {
	// Add add a new item
	Add(d *apmodelsv1.Book) (string, error)
}
