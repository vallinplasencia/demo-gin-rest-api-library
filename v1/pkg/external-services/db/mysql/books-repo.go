package mysql

import (
	"strconv"

	"database/sql"

	apmodelsv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

// booksRepo db access
type booksRepo struct {
	db *sql.DB
}

// Add add a new item
func (r *booksRepo) Add(d *apmodelsv1.Book) (string, error) {
	result, e := r.db.Exec("INSERT INTO books (title,original,tags,published_at,created_at,updated_at,category_id,user_id) VALUES (?,?,?,?,?,?,?,?)", d.Title, d.Original, d.Tags, d.PublishedAt, d.CreatedAt, d.UpdatedAt, d.CategoryID, d.UserID)
	if e != nil {
		return "", e
	}
	id, e := result.LastInsertId()
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(id, 10), nil
}
