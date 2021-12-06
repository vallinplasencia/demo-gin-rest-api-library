package dgjwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	apauthabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/auth/access-token/abstract"
	apv1models "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"

	jwt "github.com/dgrijalva/jwt-go"
)

// token manega token
type token struct {
	accessTokenSecretKey string // clave secreta para general el token

	accessTokenAudience string // aud
	accessTokenIssuer   string // iss
	accessTokenLive     int64  // tiempo de vida del token

	refreshTokenSecretKey string // clave secreta para general el token de refrescar
	refreshTokenLive      int64  // tiempo de vida del refresh token

	// urlServerStoreMedias url del servidor donde se alojan los archivos subidos
	urlServerStoreMedias string

	signingMethod *jwt.SigningMethodHMAC
}

// New ...
func New(c *config) (apauthabstract.Token, error) {
	return &token{
		accessTokenLive:      c.AccessTokenLive,
		accessTokenAudience:  c.AccessTokenAudience,
		accessTokenIssuer:    c.AccessTokenIssuer,
		accessTokenSecretKey: c.AccessTokenSecretKey,

		refreshTokenLive:      c.RefreshTokenLive,
		refreshTokenSecretKey: c.RefreshTokenSecretKey,

		urlServerStoreMedias: c.URLServerStoreMedias,

		signingMethod: jwt.SigningMethodHS256,
	}, nil
}

// Create crea y retorna un token a partir de los datos del usuario
func (t *token) Create(acc *apv1models.Account) (*apauthabstract.TokenDetails, error) {
	now := time.Now().UTC().Unix()
	roles := make([]string, len(acc.Roles))
	for i, r := range acc.Roles {
		roles[i] = string(r)
	}
	u := &apauthabstract.UserClaims{
		UserID:   acc.ID,
		Fullname: acc.Fullname,
		Username: acc.Username,
		Roles:    roles,
		Avatar:   t.fullURLMedia(acc.Avatar),
		StandardClaims: &jwt.StandardClaims{
			Audience:  t.accessTokenAudience,
			ExpiresAt: now + t.accessTokenLive,
			Id:        acc.Username,
			IssuedAt:  now,
			Issuer:    t.accessTokenIssuer,
			NotBefore: now,
			Subject:   "",
		},
	}
	accToken := jwt.NewWithClaims(t.signingMethod, u)
	accTokenStr, e := accToken.SignedString([]byte(t.accessTokenSecretKey))
	if e != nil {
		return nil, e
	}
	// creando el token de refrescar
	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = u.UserID
	rtClaims["exp"] = now + t.refreshTokenLive
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshTokenStr, e := refreshToken.SignedString([]byte(t.refreshTokenSecretKey))

	return &apauthabstract.TokenDetails{
		AccessToken:  accTokenStr,
		RefreshToken: refreshTokenStr,
	}, e
}

// Decode retorna los datos del usuario a partir de un token
func (t *token) Decode(tokenStr string) (*apauthabstract.UserClaims, error) {
	token, e := jwt.ParseWithClaims(tokenStr, &apauthabstract.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			alg := fmt.Sprint(token.Header["alg"])
			return nil, fmt.Errorf("Metodo de firma del token inesperada(%s)", alg)
		}
		return []byte(t.accessTokenSecretKey), nil
	})
	if e == nil {
		if claims, ok := token.Claims.(*apauthabstract.UserClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("Invalid token")
	}
	return nil, e
}

// DecodeYetInvalid retorna los datos del usuario a partir de un token aun estando el token invalido.
func (t *token) DecodeYetInvalid(tokenStr string) (*apauthabstract.UserClaims, error) {
	token, e := jwt.ParseWithClaims(tokenStr, &apauthabstract.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			alg := fmt.Sprint(token.Header["alg"])
			return nil, fmt.Errorf("Metodo de firma del token inesperada(%s)", alg)
		}
		return []byte(t.accessTokenSecretKey), nil
	})
	if e == nil {
		if claims, ok := token.Claims.(*apauthabstract.UserClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("invalid token")
	}
	if token != nil {
		if claims, ok := token.Claims.(*apauthabstract.UserClaims); ok {
			return claims, e
		}
	}
	// a := jwt.New(&jwt.SigningMethodHMAC{})
	// a.
	return nil, e
}

// IsErrorTokenExpired retorna true si el error es por causa de q el token ya expiro
func (t *token) IsErrorTokenExpired(e error) bool {
	v, _ := e.(*jwt.ValidationError)

	if v.Errors == jwt.ValidationErrorExpired {
		return true
	}
	return false
}

// fullURLMedia retorna la url completa de un archivo q esta almacenado en el servidor de archivos subidos
func (t *token) fullURLMedia(path string) string {
	if len(path) == 0 || strings.HasPrefix(path, "http") {
		return path
	}
	dom := strings.TrimSuffix(t.urlServerStoreMedias, "/")
	if strings.HasPrefix(path, "/") {
		return fmt.Sprintf("%s%s", dom, path)
	}
	return fmt.Sprintf("%s/%s", dom, path)
}
