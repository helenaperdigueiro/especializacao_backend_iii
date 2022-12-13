package store

// import (
	// "database/sql"
	// "avaliacao-ii/internal/domain"
// )

// type sqlStoreAppointment struct {
// 	db *sql.DB
// }

// func NewSQLStoreAppointment(db *sql.DB) StoreInterfaceAppointment {
// 	return &sqlStoreAppointment{
// 		db: db,
// 	}
// }

// func (s *sqlStoreAppointment) ReadById(id int) (domain.Appointment, error) {

// }

// func (s *sqlStoreAppointment) ReadByRg(rg string) (domain.Appointment, error) {

// }

// func (s *sqlStoreAppointment) CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {

// }

// func (s *sqlStoreAppointment) CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollment string) (domain.Appointment, error) {

// }

// func (s *sqlStoreAppointment) Update(appointment domain.Appointment) (domain.Appointment, error) {

// }

// func (s *sqlStoreAppointment) Patch(appointment domain.Appointment) (domain.Appointment, error) {

// }

// func (s *sqlStoreAppointment) Delete(id int) error {

// }