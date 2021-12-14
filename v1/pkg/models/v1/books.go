package v1

// Book ...
type Book struct {
	ID          string
	Title       string
	Original    bool
	Tags        []string
	PublishedAt int64
	CreatedAt   int64
	UpdatedAt   int64
	CategoryID  string
	UserID      string
}

// BookSimple ...
type BookSimple struct {
	ID          string
	Title       string
	Original    bool
	Tags        []string
	PublishedAt int64
	CreatedAt   int64
	UpdatedAt   int64
	UserID      string

	Category *Category
}

// BookData ...
type BookData struct {
	ID          string
	Title       string
	Original    bool
	Tags        []string
	PublishedAt int64
	CreatedAt   int64
	UpdatedAt   int64

	Category *Category
	User     *Account
}
