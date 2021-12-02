package mysql

import (
	"strconv"

	"database/sql"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	apmodelsv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

// categoriesRepo db access
type categoriesRepo struct {
	db *sql.DB
}

// Add add a new item
func (r *categoriesRepo) Add(d *apmodelsv1.Category) (string, error) {
	result, e := r.db.Exec("INSERT INTO categories (name, description, created_at,updated_at) VALUES (?,?,?,?)", d.Name, d.Description, d.CreatedAt, d.UpdatedAt)
	if e != nil {
		return "", e
	}
	id, e := result.LastInsertId()
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(id, 10), nil
}

// Find find a category by id
func (r *categoriesRepo) Find(id string) (*apmodelsv1.Category, error) {
	row := r.db.QueryRow("SELECT id,name,description,created_at,updated_at FROM categories WHERE id=?", id)
	var c apmodelsv1.Category
	switch e := row.Scan(
		&c.ID,
		&c.Name,
		&c.Description,
		&c.CreatedAt,
		&c.UpdatedAt); e {
	case nil:
		return &c, nil
	case sql.ErrNoRows:
		return nil, apdbabstract.ErrorNoItems
	default:
		return nil, e
	}
}
