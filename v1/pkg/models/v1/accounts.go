package v1

import (
	apmodels "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models"
)

// Account ...
type Account struct {
	ID        string
	Fullname  string
	Email     string
	Username  string
	Password  string
	Roles     []apmodels.RoleType
	Avatar    string
	CreatedAt int64
	UpdatedAt int64
}

// Session sessions of accessToken-RefreshToken
type Session struct {
	ID     string
	UserID string

	RefreshToken string

	DeviceID     string
	UserAgentStr string
	UserAgent    string
	Platform     string
	IP           string
	Location     string

	// LastAccessTokenGeneratedAt ultimo access token  se genero con este refresh token
	LastAccessTokenGeneratedAt int64
	CreatedAt                  int64
}
