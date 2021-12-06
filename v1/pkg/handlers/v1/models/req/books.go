package req

// CreateBook ...
type CreateBook struct {
	Title       string   `json:"title" binding:"required"`
	Original    bool     `json:"original"`
	Tags        []string `json:"tags" binding:"required"`
	PublishedAt int64    `json:"published_at"`
	CategoryID  string   `json:"category_id" binding:"required"`
}