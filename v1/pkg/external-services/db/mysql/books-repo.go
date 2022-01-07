package mysql

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"database/sql"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	apmodelsv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models"
)

// booksRepo db access
type booksRepo struct {
	db *sql.DB
}

// Add add a new item
func (r *booksRepo) Add(d *apmodelsv1.Book) (string, error) {
	jsonTags, e := json.Marshal(d.Tags)
	if e != nil {
		return "", e
	}
	q := fmt.Sprintf("INSERT INTO %s (title,original,tags,published_at,created_at,updated_at,category_id,user_id) VALUES (?,?,?,?,?,?,?,?)", booksTable)
	result, e := r.db.Exec(q, d.Title, d.Original, jsonTags, d.PublishedAt, d.CreatedAt, d.UpdatedAt, d.CategoryID, d.UserID)
	if e != nil {
		return "", e
	}
	id, e := result.LastInsertId()
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(id, 10), nil
}

// Edit update a item
func (r *booksRepo) Edit(d *apmodelsv1.Book) error {
	jsonTags, e := json.Marshal(d.Tags)
	if e != nil {
		return e
	}
	q := fmt.Sprintf("UPDATE %s SET title=?,original=?,tags=?,published_at=?,category_id=? WHERE id=?", booksTable)
	_, e = r.db.Exec(q, d.Title, d.Original, jsonTags, d.PublishedAt, d.CategoryID, d.ID)
	if e != nil {
		return e
	}
	return nil
}

// Remove delete item by id
func (r *booksRepo) Remove(id string) error {
	q := fmt.Sprintf("DELETE FROM %s WHERE id=?", booksTable)
	_, e := r.db.Exec(q, id)
	return e
}

// Find find a book by id
func (r *booksRepo) Find(id string) (*apmodelsv1.Book, error) {
	q := fmt.Sprintf(`SELECT 
	id,title,original,tags,published_at,created_at,updated_at,category_id,user_id 
	FROM %s WHERE id=?`, booksTable)
	row := r.db.QueryRow(q, id)

	item := apmodelsv1.Book{
		ID:          id,
		Title:       "",
		Original:    false,
		Tags:        []string{},
		PublishedAt: 0,
		CreatedAt:   0,
		UpdatedAt:   0,
		UserID:      "",
		CategoryID:  "",
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
		&item.CategoryID,
		&item.UserID); e {
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
		UserID:      "",
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

// ListBooksSimple list books with format simple
func (r *booksRepo) ListBooksSimple(page, limit uint, s *apdbabstract.SortBooks, f *apdbabstract.FilterBooks) ([]*apmodelsv1.BookSimple, error) {
	items := []*apmodelsv1.BookSimple{}
	// page tiene q ser mayor q 0
	offset := (page - 1) * limit
	// ordenar resultados
	orderBy := "title"
	sort := "ASC"
	// orden de los resultados
	if s != nil {
		switch s.Field {
		case apdbabstract.BookFieldTitle:
			orderBy = "title"
		case apdbabstract.BookFieldOriginal:
			orderBy = "original"
		case apdbabstract.BookFieldPublishedAt:
			orderBy = "published_at"
		case apdbabstract.BookFieldCategoryName:
			orderBy = "c.name"
		}
		if s.Descendent {
			sort = "DESC"
		}
	}
	params := []interface{}{}
	conditions := []string{}
	clauseFilterWhere := ""
	if f != nil {
		for _, values := range f.Items {
			for _, v := range values {

				switch v.Field {
				case apdbabstract.BookFieldTitle:
					cond, v, e := toConditionMysql("title", v.Values, valueString, v.Operator)
					if e != nil {
						return items, e
					}
					conditions = append(conditions, cond)
					params = append(params, v)
				case apdbabstract.BookFieldOriginal:
					cond, v, e := toConditionMysql("original", v.Values, valueBoolen, v.Operator)
					if e != nil {
						return items, e
					}
					conditions = append(conditions, cond)
					params = append(params, v)
				case apdbabstract.BookFieldTags:
					cond, v, e := toConditionMysql("tags", v.Values, valueString, v.Operator)
					if e != nil {
						return items, e
					}
					conditions = append(conditions, cond)
					params = append(params, v)
				}
			}
		}
		if len(conditions) > 0 {
			clauseFilterWhere = fmt.Sprintf("%s %s", "WHERE", strings.Join(conditions, " AND "))
		}
	}
	q := fmt.Sprintf(`SELECT
	b.id as bid,b.title as btitle,b.original as boriginal,b.tags as btags,b.published_at as bpublishedat,b.created_at as bcreatedat,b.updated_at as bupdatedat,b.user_id as buserid,c.id as cid,c.name as cname,c.description as cdescription,c.created_at as ccreateat,c.updated_at as cupdatedat
	FROM %s b INNER JOIN %s c
	ON b.category_id=c.id 
	%s 
	ORDER BY %s %s LIMIT %d,%d`, booksTable, categoriesTable, clauseFilterWhere, orderBy, sort, offset, limit)
	fmt.Println(q)
	result, e := r.db.Query(q, params...)

	if e != nil {
		return items, e
	}
	for result.Next() {
		item := apmodelsv1.BookSimple{
			ID:          "",
			Title:       "",
			Original:    false,
			Tags:        []string{},
			PublishedAt: 0,
			CreatedAt:   0,
			UpdatedAt:   0,
			UserID:      "",
			Category:    &apmodelsv1.Category{},
		}
		bitOriginal := []byte{0}
		jsonTags := ""
		e = result.Scan(
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
			&item.Category.UpdatedAt)
		if e != nil {
			return items, e
		}
		item.Original = len(bitOriginal) > 0 && bitOriginal[0] == 1
		json.Unmarshal([]byte(jsonTags), &item.Tags)
		items = append(items, &item)
	}
	return items, nil
}
