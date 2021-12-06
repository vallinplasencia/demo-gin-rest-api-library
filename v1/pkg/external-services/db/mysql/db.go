package mysql

import (
	// database/sql"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
)

const accountsTable string = "accounts"
const sessionsTable string = "sessions"
const categoriesTable string = "categories"
const booksTable string = "books"

type DB struct {
	books      apdbabstract.BooksRepo
	categories apdbabstract.CategoriesRepo
	accounts   apdbabstract.AccountsRepo
	sessions   apdbabstract.SessionsRepo
}

func New(c *config) (apdbabstract.DB, error) {
	// cc := mysql.Config{
	// 	User:                 c.User,
	// 	Passwd:               c.Password,
	// 	Net:                  "tcp",
	// 	Addr:                 c.Address, // "127.0.0.1:3306"
	// 	DBName:               c.DBName,
	// 	AllowNativePasswords: true,
	// }
	cc := mysql.NewConfig()
	cc.User = c.User
	cc.Passwd = c.Password
	cc.Addr = c.Address // "127.0.0.1:3306"
	cc.DBName = c.DBName

	dba, e := sql.Open("mysql", cc.FormatDSN())
	if e != nil {
		return nil, e
	}
	dba.SetConnMaxLifetime(time.Minute * 3)
	dba.SetMaxOpenConns(10)
	dba.SetMaxIdleConns(10)

	if e = dba.Ping(); e != nil {
		return nil, e
	}
	rBooks := &booksRepo{
		db: dba,
	}
	rCategories := &categoriesRepo{
		db: dba,
	}
	rAccounts := &accountsRepo{
		db: dba,
	}
	rSessions := &sessionsRepo{
		db: dba,
	}
	db := &DB{
		books:      rBooks,
		categories: rCategories,
		accounts:   rAccounts,
		sessions:   rSessions,
	}
	return db, nil
}

func (db *DB) Books() apdbabstract.BooksRepo {
	return db.books
}
func (db *DB) Categories() apdbabstract.CategoriesRepo {
	return db.categories
}
func (db *DB) Accounts() apdbabstract.AccountsRepo {
	return db.accounts
}
func (db *DB) Sessions() apdbabstract.SessionsRepo {
	return db.sessions
}
