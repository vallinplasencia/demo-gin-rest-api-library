package models

const (
	// KeyUserContext clave en el contexto de gin para obtener el usuario logueado en la app
	KeyUserContext string = "user_in"
)

// AuthUser datos del usuario logueado y q se obtiene a traves del contexto de gin con esta clave KeyUserContext
type AuthUser struct {
	UserID      string
	Fullname    string
	Username    string
	Roles       []RoleType
	Permissions map[PermissionType]bool
	Avatar      string
}

// IsAnonymous retorna true si el usuario no esta logueado
func (u *AuthUser) IsAnonymous() bool {
	for _, v := range u.Roles {
		if v == RoleAnonymous {
			return true
		}
	}
	return false
}

// ContainRermission retorna true si el usuario tiene el permiso p
func (u *AuthUser) ContainRermission(p PermissionType) bool {
	_, ok := u.Permissions[p]
	return ok
}

// === types === //

// RoleType users role
type RoleType string

const (
	RoleAnonymous RoleType = "anonymous"
	RoleUser      RoleType = "user"
	RoleAdmin     RoleType = "admin"
)

// PermissionType permisos asignados a usuarios segun su rol
type PermissionType int

const (
	// === accounts === //
	PermissionCreateAccount PermissionType = iota + 1
	PermissionLogin

	// === books === //
	PermissionAddBook
)

// ToRoleFromString retorna un tipo de rol segun s.
func ToRoleFromString(s string) RoleType {
	values := map[string]RoleType{
		"anonymous": RoleAnonymous,
		"user":      RoleUser,
		"admin":     RoleAdmin,
	}
	r, ok := values[s]
	if ok {
		return r
	}
	return RoleAnonymous
}
