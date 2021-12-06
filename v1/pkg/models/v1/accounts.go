package v1

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

type RoleType string

const (
	RoleIndeterminate RoleType = ""
	RoleAnonymous     RoleType = "anonymous"
	RoleUser          RoleType = "user"
	RoleAdmin         RoleType = "admin"
)

// ToRolesFromString retorna un tipo de rol segun s.
//
func ToRolesFromString(s string) RoleType {
	values := map[string]RoleType{
		"":          RoleIndeterminate,
		"anonymous": RoleAnonymous,
		"user":      RoleUser,
		"admin":     RoleAdmin,
	}
	r, ok := values[s]
	if ok {
		return r
	}
	return RoleIndeterminate
}
