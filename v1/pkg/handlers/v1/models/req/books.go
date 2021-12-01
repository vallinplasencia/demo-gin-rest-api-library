package req

// CreateBook ...
type CreateBook struct {
	Title       string `json:"title" binding:"required"`
	Original    bool   `json:"original"`
	CategoryID  string `json:"category_id" binding:"required"`
	PublishedAt int64  `json:"published_at"`
}
