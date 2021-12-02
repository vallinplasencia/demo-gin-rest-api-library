package mysql

import (
	"strconv"

	"database/sql"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	apmodelsv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

// accountsRepo db access
type accountsRepo struct {
	db *sql.DB
}

// Add add a new item
func (r *accountsRepo) Add(d *apmodelsv1.Account) (string, error) {
	result, e := r.db.Exec("INSERT INTO accounts (fullname,email,username,password,created_at,updated_at) VALUES (?,?,?,?,?,?)", d.Fullname, d.Email, d.Username, d.Password, d.CreatedAt, d.UpdatedAt)
	if e != nil {
		return "", e
	}
	id, e := result.LastInsertId()
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(id, 10), nil
}

// Find find a account by id
func (r *accountsRepo) Find(id string) (*apmodelsv1.Account, error) {
	row := r.db.QueryRow("SELECT id,fullname,email,username,password,created_at,updated_at FROM accounts WHERE id=?", id)
	var d apmodelsv1.Account
	switch e := row.Scan(
		&d.ID,
		&d.Fullname,
		&d.Email,
		&d.Username,
		&d.Password,
		&d.CreatedAt,
		&d.UpdatedAt); e {
	case nil:
		return &d, nil
	case sql.ErrNoRows:
		return nil, apdbabstract.ErrorNoItems
	default:
		return nil, e
	}
}

// FindByUsername find a account by username
func (r *accountsRepo) FindByUsername(username string) (*apmodelsv1.Account, error) {
	row := r.db.QueryRow("SELECT id,fullname,email,username,password,created_at,updated_at FROM accounts WHERE username=?", username)
	var d apmodelsv1.Account
	switch e := row.Scan(
		&d.ID,
		&d.Fullname,
		&d.Email,
		&d.Username,
		&d.Password,
		&d.CreatedAt,
		&d.UpdatedAt); e {
	case nil:
		return &d, nil
	case sql.ErrNoRows:
		return nil, apdbabstract.ErrorNoItems
	default:
		return nil, e
	}
}

// FindByEmail find a account by email
func (r *accountsRepo) FindByEmail(email string) (*apmodelsv1.Account, error) {
	row := r.db.QueryRow("SELECT id,fullname,email,username,password,created_at,updated_at FROM accounts WHERE email=?", email)
	var d apmodelsv1.Account
	switch e := row.Scan(
		&d.ID,
		&d.Fullname,
		&d.Email,
		&d.Username,
		&d.Password,
		&d.CreatedAt,
		&d.UpdatedAt); e {
	case nil:
		return &d, nil
	case sql.ErrNoRows:
		return nil, apdbabstract.ErrorNoItems
	default:
		return nil, e
	}
}
