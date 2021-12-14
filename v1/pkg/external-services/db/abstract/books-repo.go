package abstract

import (
	apmodelsv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

// BooksRepo db access
type BooksRepo interface {
	// Add add a new item
	Add(d *apmodelsv1.Book) (string, error)
	// Edit update a item
	Edit(d *apmodelsv1.Book) error

	// FindBook find a book by id
	FindBook(id string) (*apmodelsv1.Book, error)
	// FindBookSimple find a book by id con su categoria
	FindBookSimple(id string) (*apmodelsv1.BookSimple, error)
}
