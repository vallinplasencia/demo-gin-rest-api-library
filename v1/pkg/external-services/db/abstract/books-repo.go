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
	Find(id string) (*apmodelsv1.Book, error)
	// Remove delete item by id
	Remove(id string) error

	// FindBookSimple find a book by id con su categoria
	FindBookSimple(id string) (*apmodelsv1.BookSimple, error)

	// ListBooksSimple list books with format simple
	ListBooksSimple(page, limit uint, s *SortBooks, f *FilterBooks) ([]*apmodelsv1.BookSimple, error)
}

// BookFieldType field for book. Util for filters, sort, ...
type BookFieldType int

const (
	// BookFieldIndeterminate BookFieldType = iota
	// BookFieldTitle ..
	BookFieldTitle BookFieldType = iota
	// BookFieldOriginal ...
	BookFieldOriginal
	// // BookFieldTag ...
	// BookFieldTag
	// BookFieldTags ...
	BookFieldTags
	// BookFieldPublishedAt ...
	BookFieldPublishedAt
	// BookFieldCategoryName ...
	BookFieldCategoryName
	// BookFieldCategoryID ...
	BookFieldCategoryID
	// BookFieldUserID ...
	BookFieldUserID
)

// SortBooks sort list books
type SortBooks struct {
	Field      BookFieldType
	Descendent bool
}

// FilterBooks filter list books
type FilterBooks struct {
	Items map[BookFieldType][]*FilterBooksItem
}

// FilterBooksItem filtro para un campo
type FilterBooksItem struct {
	Field    BookFieldType
	Value    string
	Values   []string
	Operator OperatorType
}
