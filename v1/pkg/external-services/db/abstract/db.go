package abstract

// DB acceso a la db
type DB interface {
	Books() BooksRepo
	Categories() CategoriesRepo
	Accounts() AccountsRepo
	Sessions() SessionsRepo
}