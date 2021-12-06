package mysql

import (
	"fmt"
	"strconv"

	"database/sql"

	apmodelsv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

const tableName string = "sessions"

// sessionsRepo db access
type sessionsRepo struct {
	db *sql.DB
}

// Add add a new item
func (r *sessionsRepo) Add(d *apmodelsv1.Session) (string, error) {
	q := fmt.Sprintf("INSERT INTO %s (refresh_token,device_id,useragent_str,useragent,platform,ip,location,last_access_token_generated_at,created_at,user_id) VALUES (?,?,?,?,?,?,?,?,?,?)", tableName)
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
