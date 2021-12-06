package mysql

import (
	"fmt"
	"strconv"
	"strings"

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
	roles := make([]string, len(d.Roles))
	for i, v := range d.Roles {
		roles[i] = string(v)
	}
	strRoles := fmt.Sprintf("%s", strings.Join(roles, ","))
	// strRoles = "user"
	fmt.Println("AAAAA: ", strRoles)
	q := fmt.Sprintf("INSERT INTO %s (fullname,email,username,password,roles,avatar,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?)", accountsTable)
	result, e := r.db.Exec(q, d.Fullname, d.Email, d.Username, d.Password, strRoles, d.Avatar, d.CreatedAt, d.UpdatedAt)
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
	q := fmt.Sprintf("SELECT id,fullname,email,username,password,roles,avatar,created_at,updated_at FROM %s WHERE id=?", accountsTable)
	row := r.db.QueryRow(q, id)
	roles := ""
	var d apmodelsv1.Account
	switch e := row.Scan(
		&d.ID,
		&d.Fullname,
		&d.Email,
		&d.Username,
		&d.Password,
		&roles,
		&d.Avatar,
		&d.CreatedAt,
		&d.UpdatedAt); e {
	case nil:
		d.Roles = r.toRoles(roles)
		return &d, nil
	case sql.ErrNoRows:
		return nil, apdbabstract.ErrorNoItems
	default:
		return nil, e
	}
}

// FindByUsername find a account by username
func (r *accountsRepo) FindByUsername(username string) (*apmodelsv1.Account, error) {
	q := fmt.Sprintf("SELECT id,fullname,email,username,password,roles,avatar,created_at,updated_at FROM %s WHERE username=?", accountsTable)
	row := r.db.QueryRow(q, username)
	roles := ""
	var d apmodelsv1.Account
	switch e := row.Scan(
		&d.ID,
		&d.Fullname,
		&d.Email,
		&d.Username,
		&d.Password,
		&roles,
		&d.Avatar,
		&d.CreatedAt,
		&d.UpdatedAt); e {
	case nil:
		d.Roles = r.toRoles(roles)
		return &d, nil
	case sql.ErrNoRows:
		return nil, apdbabstract.ErrorNoItems
	default:
		return nil, e
	}
}

// FindByEmail find a account by email
func (r *accountsRepo) FindByEmail(email string) (*apmodelsv1.Account, error) {
	q := fmt.Sprintf("SELECT id,fullname,email,username,password,roles,avatar,created_at,updated_at FROM %s WHERE email=?", accountsTable)
	row := r.db.QueryRow(q, email)
	roles := ""
	var d apmodelsv1.Account
	switch e := row.Scan(
		&d.ID,
		&d.Fullname,
		&d.Email,
		&d.Username,
		&d.Password,
		&roles,
		&d.Avatar,
		&d.CreatedAt,
		&d.UpdatedAt); e {
	case nil:
		d.Roles = r.toRoles(roles)
		return &d, nil
	case sql.ErrNoRows:
		return nil, apdbabstract.ErrorNoItems
	default:
		return nil, e
	}
}

func (r *accountsRepo) toRoles(strRoles string) []apmodelsv1.RoleType {
	roles := []apmodelsv1.RoleType{}
	for _, v := range strings.Split(strRoles, ",") {
		roles = append(roles, apmodelsv1.ToRolesFromString(v))
	}
	return roles
}
