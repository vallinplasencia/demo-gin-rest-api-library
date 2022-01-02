package req

// CreateBook ...
type CreateBook struct {
	Title       string   `json:"title" binding:"required"`
	Original    bool     `json:"original"`
	Tags        []string `json:"tags" binding:"required"`
	PublishedAt int64    `json:"published_at"`
	CategoryID  string   `json:"category_id" binding:"required"`
}

// EditBook ...
type EditBook struct {
	Title       string   `json:"title" binding:"required"`
	Original    bool     `json:"original"`
	Tags        []string `json:"tags" binding:"required"`
	PublishedAt int64    `json:"published_at"`
	CategoryID  string   `json:"category_id" binding:"required"`
}

// // === sort === //

// // SortListBooks sort list books
// type SortListBooks struct {
// 	By         BookFieldType
// 	Descendent bool
// }

// // BookFieldType field for book. Util for filters, sort, ...
// type BookFieldType int

// const (
// 	BookFieldIndeterminate BookFieldType = iota
// 	// BookFieldTitle ..
// 	BookFieldTitle
// 	// BookFieldOriginal ...
// 	BookFieldOriginal
// 	// BookFieldTags ...
// 	BookFieldTags
// 	// BookFieldPublishedAt ...
// 	BookFieldPublishedAt
// 	// CategoryID ...
// 	CategoryID
// )

// // ToBookFieldFromRequest convert of string field to BookFieldType.
// //
// // action validate field for sort, filter, ...
// //
// // if field not found return BookFieldIndeterminate
// func ToBookFieldFromRequest(field string) BookFieldType {
// 	fields := map[string]BookFieldType{
// 		"title":        BookFieldTitle,
// 		"original":     BookFieldOriginal,
// 		"tags":         BookFieldTags,
// 		"published-at": BookFieldPublishedAt,
// 		"category-id":  CategoryID,
// 	}
// 	v, ok := fields[field]
// 	if ok {
// 		return v
// 	}
// 	return BookFieldIndeterminate
// }
