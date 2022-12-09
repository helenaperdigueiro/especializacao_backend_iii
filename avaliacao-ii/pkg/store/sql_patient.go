package store

import (
	"database/sql"
	"avaliacao-ii/internal/domain"
)

type sqlStorePatient struct {
	db *sql.DB
}

func NewSQLStorePatient(db *sql.DB) StoreInterfacePatient {
	return &sqlStorePatient{
		db: db,
	}
}

func (s *sqlStorePatient) ReadById(id int) (domain.Patient, error) {

}
func (s *sqlStorePatient) Create(patient domain.Patient) (domain.Patient, error) {

}
func (s *sqlStorePatient) Update(patient domain.Patient) (domain.Patient, error) {

}
func (s *sqlStorePatient) Patch(patient domain.Patient) (domain.Patient, error) {

}
func (s *sqlStorePatient) Delete(id int) error {

}