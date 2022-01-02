package abstract

// DB acceso a la db
type DB interface {
	Books() BooksRepo
	Categories() CategoriesRepo
	Accounts() AccountsRepo
	Sessions() SessionsRepo
}

// OperatorType operadores q se pueden usar en consultas a bd
type OperatorType int

const (
	// OperatorEqual ...
	OperatorEqual OperatorType = iota
	// OperatorNotEqual ...
	OperatorNotEqual
	// OperatorLessThan ...
	OperatorLessThan
	// OperatorGreatThan ...
	OperatorGreatThan
	// OperatorLessThanEqual ...
	OperatorLessThanEqual
	// OperatorGreatThanEqual ...
	OperatorGreatThanEqual
	// OperatorRange ...
	OperatorRange
	// OperatorStartWith ...
	OperatorStartWith
	// OperatorEndWith ...
	OperatorEndWith
	// OperatorContain ...
	OperatorContain
)
