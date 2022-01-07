package models

// Account ...
type Account struct {
	ID        string
	Fullname  string
	Email     string
	Username  string
	Password  string
	Roles     []RoleType
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
