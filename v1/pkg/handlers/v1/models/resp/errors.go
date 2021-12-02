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
	CodeConflictEmail
	CodeConflictUsername
	CodeNotFoundUser
	CodeIncorrectPassword
)

var msgs = map[CodeType]string{
	CodeIndeterminate:     "",
	CodeOK:                "ok",
	CodeInternalError:     "internal error",
	CodeInvalidArgument:   "invalid arguments",
	CodeNotFoundCategory:  "category not found",
	CodeConflictEmail:     "email exist",
	CodeConflictUsername:  "username exist",
	CodeNotFoundUser:      "user not found",
	CodeIncorrectPassword: "password incorrect",
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
