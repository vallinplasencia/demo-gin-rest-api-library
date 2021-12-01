package mysql

import (
	// database/sql"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
)

type DB struct {
	books      apdbabstract.BooksRepo
	categories apdbabstract.CategoriesRepo
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
	db := &DB{
		books:      rBooks,
		categories: rCategories,
	}
	return db, nil
}

func (db *DB) Books() apdbabstract.BooksRepo {
	return db.books
}
func (db *DB) Categories() apdbabstract.CategoriesRepo {
	return db.categories
}
