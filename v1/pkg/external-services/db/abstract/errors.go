package abstract

import "errors"

var (
	// ErrNoItems ...
	ErrorNoItems = errors.New("no items")
	// ErrorOperatorInvalidForValueType cuando se intenta aplicar un operador a un valor cuyo tipo no es permitido
	ErrorOperatorInvalidForValueType = errors.New("operator is invalid for value type")
	// ErrorOperatorInvalid operator invalid
	ErrorOperatorInvalid = errors.New("operator is invalid")
)
