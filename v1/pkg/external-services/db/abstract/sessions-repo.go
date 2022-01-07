package abstract

import (
	apmodelsv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models"
)

// SessionsRepo db access
type SessionsRepo interface {
	// Add add a new item
	Add(d *apmodelsv1.Session) (string, error)
	// Remove delete item by id
	Remove(id string) error
	// EditLastAccessTokenGenerated update lastAccessTokenGeneratedAt
	EditLastAccessTokenGenerated(id string, lastAccessTokenGeneratedAt int64) error
	// Find find a session by refresh token
	FindByRefreshToken(refreshToken string) (*apmodelsv1.Session, error)
}
