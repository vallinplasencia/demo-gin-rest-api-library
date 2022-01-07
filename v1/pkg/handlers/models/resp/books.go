package resp

// BookSimple ...
type BookSimple struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Original    bool     `json:"original"`
	Tags        []string `json:"tags"`
	PublishedAt int64    `json:"publishet_at"`
	CreatedAt   int64    `json:"created_at"`
	UpdatedAt   int64    `json:"updated_at"`
	UserID      string   `json:"user_id"`

	Category *CategorySimple `json:"category"`
}
