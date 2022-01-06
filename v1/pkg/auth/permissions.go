package auth

import (
	apmodels "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models"
)

// permissions permisos asignados a cada rol de usuario
var permissions = map[apmodels.RoleType][]apmodels.PermissionType{}

// GetPermissions return all permissions by roles
func GetPermissions(roles []apmodels.RoleType) map[apmodels.PermissionType]bool {
	perms := map[apmodels.PermissionType]bool{}
	for _, v := range roles {
		for _, v := range permissions[v] {
			perms[v] = true
		}
	}
	return perms
}

func init() {
	// === carga los permisos por roles === //

	anonymous := []apmodels.PermissionType{
		apmodels.PermissionCreateAccount, apmodels.PermissionLogin, apmodels.PermissionGenerateAccessToken,
	}
	users := []apmodels.PermissionType{
		apmodels.PermissionAddBook, apmodels.PermissionRetrieveBook, apmodels.PermissionEditBook, apmodels.PermissionListBooks, apmodels.PermissionGenerateAccessToken,
	}
	admin := []apmodels.PermissionType{}
	permissions = map[apmodels.RoleType][]apmodels.PermissionType{
		apmodels.RoleUser:      users,
		apmodels.RoleAnonymous: anonymous,
		apmodels.RoleAdmin:     admin,
	}
}
