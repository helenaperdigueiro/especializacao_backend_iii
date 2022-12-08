package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type sqlStorePatient struct {
	db *sql.DB
}

func NewSQLStorePatient(db *sql.DB) StoreInterfacePatient {
	return &sqlStorePatient{
		db: db,
	}
}