package mysql

import (
	"encoding/json"
	"fmt"
	"strconv"

	"database/sql"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	apmodelsv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

// booksRepo db access
type booksRepo struct {
	db *sql.DB
}

// Add add a new item
func (r *booksRepo) Add(d *apmodelsv1.Book) (string, error) {
	// tags := make([]string, len(d.Tags))
	// for i, v := range d.Tags {
	// 	tags[i] = fmt.Sprintf("\"%s\"", v)
	// }
	// strTags := fmt.Sprintf("[%s]", strings.Join(tags, ","))
	jsonTags, e := json.Marshal(d.Tags)
	if e != nil {
		return "", e
	}
	result, e := r.db.Exec("INSERT INTO books (title,original,tags,published_at,created_at,updated_at,category_id,user_id) VALUES (?,?,?,?,?,?,?,?)", d.Title, d.Original, jsonTags, d.PublishedAt, d.CreatedAt, d.UpdatedAt, d.CategoryID, d.UserID)
	if e != nil {
		return "", e
	}
	id, e := result.LastInsertId()
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(id, 10), nil
}

// FindBookSimple find a book by id con su categoria
func (r *booksRepo) FindBookSimple(id string) (*apmodelsv1.BookSimple, error) {
	q := fmt.Sprintf(`SELECT 
	b.id as bid,b.title as btitle,b.original as boriginal,b.tags as btags,b.published_at as bpublishedat,b.created_at as bcreatedat,b.updated_at as bupdatedat,b.user_id as buserid,c.id as cid,c.name as cname,c.description as cdescription,c.created_at as ccreateat,c.updated_at as cupdatedat 
	FROM %s b INNER JOIN %s c 
	ON b.category_id=c.id AND b.id=?`, booksTable, categoriesTable)
	row := r.db.QueryRow(q, id)

	item := apmodelsv1.BookSimple{
		ID:          id,
		Title:       "",
		Original:    false,
		Tags:        []string{},
		PublishedAt: 0,
		CreatedAt:   0,
		UpdatedAt:   0,
		UserID:      id,
		Category:    &apmodelsv1.Category{},
	}

	bitOriginal := []byte{0}
	jsonTags := ""
	switch e := row.Scan(
		&item.ID,
		&item.Title,
		&bitOriginal,
		&jsonTags,
		&item.PublishedAt,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.UserID,
		&item.Category.ID,
		&item.Category.Name,
		&item.Category.Description,
		&item.Category.CreatedAt,
		&item.Category.UpdatedAt); e {
	case nil:
		item.Original = len(bitOriginal) > 0 && bitOriginal[0] == 1
		json.Unmarshal([]byte(jsonTags), &item.Tags)
		return &item, nil
	case sql.ErrNoRows:
		return nil, apdbabstract.ErrorNoItems
	default:
		return nil, e
	}
}
