package store

import (
	"database/sql"
	"avaliacao-ii/internal/domain"
)

type sqlStoreDentist struct {
	db *sql.DB
}

func NewSQLStoreDentist(db *sql.DB) StoreInterfaceDentist {
	return &sqlStoreDentist{
		db: db,
	}
}

func (s *sqlStoreDentist) ReadById(id int) (domain.Dentist, error) {

}
func (s *sqlStoreDentist) Create(dentist domain.Dentist) (domain.Dentist, error) {

}
func (s *sqlStoreDentist) Update(dentist domain.Dentist) (domain.Dentist, error) {

}
func (s *sqlStoreDentist) Patch(dentist domain.Dentist) (domain.Dentist, error) {

}
func (s *sqlStoreDentist) Delete(id int) error {

}