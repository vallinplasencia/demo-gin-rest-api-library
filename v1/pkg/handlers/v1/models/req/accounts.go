package req

import "mime/multipart"

// CreateAccount ...
type CreateAccount struct {
	Fullname string                `form:"fullname" binding:"required"`
	Email    string                `form:"email" binding:"required"`
	Password string                `form:"password" binding:"required"`
	Gender   GenderType            `form:"gender" binding:"required"`
	Avatar   *multipart.FileHeader `form:"avatar" binding:"required"`
}

// // CreateAccount ...
// type CreateAccount struct {
// 	Fullname string     `json:"fullname" binding:"required"`
// 	Email    string     `json:"email" binding:"required"`
// 	Username string     `json:"username" binding:"required"`
// 	Password string     `json:"password" binding:"required"`
// 	Gender   GenderType `json:"gender" binding:"required"`
// 	Avatar   string     `json:"avatar"`
// }

// Login ...
type Login struct {
	UsernameOrEmail string `json:"username_or_email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

// === types === //

// GenderType gender type
type GenderType string

const (
	// GenderIndeterminate genero incorrecto o son valor
	GenderIndeterminate GenderType = ""
	// GenderMale ...
	GenderMale GenderType = "male"
	// GenderFemale...
	GenderFemale GenderType = "female"
	// GenderOther ...
	GenderOther GenderType = "other"
)
