package abstract

import (
	"errors"

	"github.com/dgrijalva/jwt-go"

	apmodels "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models"
)

var (
	// ErrUnexpectedSigningMethod metodo de firma del token no coincide con el q se usa
	ErrUnexpectedSigningMethod error = errors.New("Unexpected signing method")
	// ErrInvalidToken ...
	ErrInvalidToken error = errors.New("invalid token")
)

// Token ...
type Token interface {
	// Create crea y retorna un token a partir de los datos del usuario
	Create(u *apmodels.AuthUser) (*TokenDetails, error)
	// // Decode retorna los datos del usuario a partir de un token
	Decode(token string) (*UserClaims, error)
	// DecodeYetInvalid retorna los datos del usuario a partir de un token aunque el token sea invalido.
	DecodeYetInvalid(tokenStr string) (*UserClaims, error)
	// GetLiveRefreshToken retorna la cantidad de segundos q es valido el refresh token
	GetLiveRefreshToken() int64
}

// UserClaims datos q vienen en el token
type UserClaims struct {
	UserID   string   `json:"user_id"`
	Fullname string   `json:"fullname"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	Avatar   string   `json:"avatar"`

	*jwt.StandardClaims
}

//  TokenDetails ...
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
}
