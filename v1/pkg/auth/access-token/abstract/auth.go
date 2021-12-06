package abstract

import (
	"github.com/dgrijalva/jwt-go"

	apv1models "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

// Token ...
type Token interface {
	// Create crea y retorna un token a partir de los datos del usuario
	Create(u *apv1models.Account) (*TokenDetails, error)
	// // Decode retorna los datos del usuario a partir de un token
	// Decode(token string) (*UserClaims, error)
	// // DecodeYetInvalid retorna los datos del usuario a partir de un token aun estando el token invalido.
	// DecodeYetInvalid(tokenStr string) (*UserClaims, error)
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
