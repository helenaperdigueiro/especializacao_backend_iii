package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type sqlStoreAppointment struct {
	db *sql.DB
}

func NewSQLStoreAppointment(db *sql.DB) StoreInterfaceAppointment {
	return &sqlStoreAppointment{
		db: db,
	}
}