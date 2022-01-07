package mysql

import (
	"fmt"
	"strconv"

	"database/sql"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	apmodelsv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models"
)

// sessionsRepo db access
type sessionsRepo struct {
	db *sql.DB
}

// Add add a new item
func (r *sessionsRepo) Add(d *apmodelsv1.Session) (string, error) {
	q := fmt.Sprintf("INSERT INTO %s (refresh_token,device_id,useragent_str,useragent,platform,ip,location,last_access_token_generated_at,created_at,user_id) VALUES (?,?,?,?,?,?,?,?,?,?)", sessionsTable)
	result, e := r.db.Exec(q, d.RefreshToken, d.DeviceID, d.UserAgentStr, d.UserAgent, d.Platform, d.ID, d.Location, d.LastAccessTokenGeneratedAt, d.CreatedAt, d.UserID)
	if e != nil {
		return "", e
	}
	id, e := result.LastInsertId()
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(id, 10), nil
}

// Remove delete item by id
func (r *sessionsRepo) Remove(id string) error {
	q := fmt.Sprintf("DELETE FROM %s WHERE id=?", sessionsTable)
	_, e := r.db.Exec(q, id)
	return e
}

// EditLastAccessTokenGenerated update lastAccessTokenGeneratedAt
func (r *sessionsRepo) EditLastAccessTokenGenerated(id string, lastAccessTokenGeneratedAt int64) error {
	q := fmt.Sprintf("UPDATE %s SET last_access_token_generated_at=? WHERE id=?", sessionsTable)
	_, e := r.db.Exec(q, lastAccessTokenGeneratedAt, id)
	if e != nil {
		return e
	}
	return nil
}

// Find find a session by refresh token
func (r *sessionsRepo) FindByRefreshToken(refreshToken string) (*apmodelsv1.Session, error) {
	q := fmt.Sprintf("SELECT id,refresh_token,device_id,useragent_str,useragent,platform,ip,location,last_access_token_generated_at,created_at,user_id FROM %s WHERE refresh_token=?", sessionsTable)
	row := r.db.QueryRow(q, refreshToken)
	var c apmodelsv1.Session
	switch e := row.Scan(
		&c.ID,
		&c.RefreshToken,
		&c.DeviceID,
		&c.UserAgentStr,
		&c.UserAgent,
		&c.Platform,
		&c.IP,
		&c.Location,
		&c.LastAccessTokenGeneratedAt,
		&c.CreatedAt,
		&c.UserID); e {
	case nil:
		return &c, nil
	case sql.ErrNoRows:
		return nil, apdbabstract.ErrorNoItems
	default:
		return nil, e
	}
}
