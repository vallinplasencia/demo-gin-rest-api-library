package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	apauth "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/auth"
	apauthtokenabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/auth/access-token/abstract"
	aphv1resp "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/models/resp"
	apmodels "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models"
)

const projectName string = "LIBRARY"

// bearerSchema esquema bearer para la cabecera authorization
const bearerSchema string = "Bearer "

// authHeader authorization header
type authHeader struct {
	BearerToken string `header:"Authorization"`
}

// AuthJwt middleware para authorizacion con un jwt(access-token)
func AuthJwt(t apauthtokenabstract.Token) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := authHeader{}
		// codigo y mensaje por defecto
		code := aphv1resp.CodeInvalidAuthToken
		msg := aphv1resp.GetMsgError(code)

		if e := c.ShouldBindHeader(&authHeader); e != nil {
			msg = fmt.Sprintf("%s --- %s", msg, e.Error())
			c.JSON(http.StatusUnauthorized, &aphv1resp.Error{
				Code: code,
				Msg:  msg,
			})
			c.Abort()
			return
		}
		if len(authHeader.BearerToken) == 0 {
			roles := []apmodels.RoleType{apmodels.RoleAnonymous}
			user := &apmodels.AuthUser{
				UserID:      "",
				Fullname:    "",
				Username:    "",
				Roles:       roles,
				Permissions: apauth.GetPermissions(roles),
				Avatar:      "",
			}
			c.Set(apmodels.KeyUserContext, user)
			c.Next()
			return
		}
		// comprobando q si la cabecera authorization viene en la peticion tenga el esquema bearer
		if !strings.HasPrefix(authHeader.BearerToken, bearerSchema) {
			msg = fmt.Sprintf("%s --- %s", msg, "token is empty or not schema bearer")
			c.JSON(http.StatusUnauthorized, &aphv1resp.Error{
				Code: code,
				Msg:  msg,
			})
			c.Abort()
			return
		}
		// decodificando el jwt(access-token)
		u, e := t.Decode(strings.TrimPrefix(authHeader.BearerToken, bearerSchema))
		// comprobando q el token sea valido
		if e != nil {
			msg = fmt.Sprintf("%s --- %s", msg, "decode token error")
			c.JSON(http.StatusUnauthorized, &aphv1resp.Error{
				Code: code,
				Msg:  msg,
			})
			c.Abort()
			return
		}
		roles := make([]apmodels.RoleType, len(u.Roles))
		for i, v := range u.Roles {
			roles[i] = apmodels.ToRoleFromString(v)
		}
		user := &apmodels.AuthUser{
			UserID:      u.UserID,
			Fullname:    u.Fullname,
			Username:    u.Username,
			Roles:       roles,
			Avatar:      u.Avatar,
			Permissions: apauth.GetPermissions(roles),
		}
		c.Set(apmodels.KeyUserContext, user)
		c.Next()
	}
}
