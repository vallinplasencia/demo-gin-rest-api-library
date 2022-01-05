package middlewares

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"

// 	apauthtokenabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/auth/access-token/abstract"
// 	aphv1resp "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1/models/resp"
// 	apmodels "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models"
// )

// // permissions permisos asignados a cada rol de usuario
// var permissions = map[apmodels.RoleType][]apmodels.PermissionType{}

// func init() {
// 	// === carga los permisos por roles === //

// 	anonymous := []apmodels.PermissionType{
// 		apmodels.PermissionCreateAccount, apmodels.PermissionLogin,
// 	}
// 	users := []apmodels.PermissionType{
// 		apmodels.PermissionAddBook, apmodels.PermissionRetrieveBook, apmodels.PermissionEditBook, apmodels.PermissionListBooks,
// 	}
// 	admin := []apmodels.PermissionType{}
// 	permissions = map[apmodels.RoleType][]apmodels.PermissionType{
// 		apmodels.RoleUser:      users,
// 		apmodels.RoleAnonymous: anonymous,
// 		apmodels.RoleAdmin:     admin,
// 	}
// }
