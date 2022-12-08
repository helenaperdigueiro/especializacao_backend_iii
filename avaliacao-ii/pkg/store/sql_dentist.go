package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type sqlStoreDentist struct {
	db *sql.DB
}

func NewSQLStoreDentist(db *sql.DB) StoreInterfaceDentist {
	return &sqlStoreDentist{
		db: db,
	}
}