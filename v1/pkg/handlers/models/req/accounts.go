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

// Login ...
type Login struct {
	UsernameOrEmail string `json:"username_or_email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

// GenerateAccessToken ...
type GenerateAccessToken struct {
	RefreshToken   string `json:"refresh_token"`
	OldAccessToken string `json:"old_access_token"`
	DeviceID       string `json:"device_id"`
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
