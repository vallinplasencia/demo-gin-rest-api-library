package resp

// Error error response
type Error struct {
	Code CodeType `json:"code"`
	Msg  string   `json:"msg"`
}

type CodeType int

const (
	CodeIndeterminate CodeType = iota + 1000
	CodeOK
	CodeInternalError
	CodeInvalidArgument
	CodeNotFoundCategory
	// CodeConflictEmail email exist
	CodeConflictEmail
	// CodeConflictUsername username exist
	CodeConflictUsername
	CodeNotFoundUser
	CodeIncorrectPassword
	// CodeInvalidAuthToken jwt invalid
	CodeInvalidAuthToken
	CodeUnauthorized
	CodeNotFoundBook
	// CodeExpiredRefreshToken expired refresh token
	CodeExpiredRefreshToken
	// CodeInvalidSessionDataOfRefreshToken session data asociated with the refresh token is invalid. Ex: diferente platform, ...
	CodeInvalidSessionDataOfRefreshToken
)

var msgs = map[CodeType]string{
	CodeIndeterminate:                    "",
	CodeOK:                               "ok",
	CodeInternalError:                    "internal error",
	CodeInvalidArgument:                  "invalid arguments",
	CodeNotFoundCategory:                 "category not found",
	CodeConflictEmail:                    "email exist",
	CodeConflictUsername:                 "username exist",
	CodeNotFoundUser:                     "user not found",
	CodeIncorrectPassword:                "password incorrect",
	CodeInvalidAuthToken:                 "invalid token auth",
	CodeUnauthorized:                     "unauthorized",
	CodeNotFoundBook:                     "book not found",
	CodeExpiredRefreshToken:              "refresh token expired",
	CodeInvalidSessionDataOfRefreshToken: "session data associated with the refresh token is invalid",
}

// GetMsgError retorna el mensaje de error asociado al codigo
//
// si no hay asociado un mensaje de error al codigo se retona un interal error mensaje
func GetMsgError(code CodeType) string {
	if msg, ok := msgs[code]; ok {
		return msg
	}
	return msgs[CodeInternalError]
}
