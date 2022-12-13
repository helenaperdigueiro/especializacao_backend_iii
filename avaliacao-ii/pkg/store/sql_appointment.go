package store

import (
	"avaliacao-ii/internal/domain"
	"database/sql"
	"errors"
)

type sqlStoreAppointment struct {
	db *sql.DB
}

func NewSQLStoreAppointment(db *sql.DB) StoreInterfaceAppointment {
	return &sqlStoreAppointment{
		db: db,
	}
}

func (s *sqlStoreAppointment) ReadById(id int) (domain.Appointment, error) {
	queryGetById := `SELECT appointment.id, patient.id, patient.name, patient.last_name, patient.rg, patient.register_date, 
							dentist.id, dentist.name, dentist.last_name, dentist.enrollment, 
							appointment.date, appointment.description 
							FROM appointment 
							INNER JOIN patient 
							ON patient.id = appointment.patient_id 
							INNER JOIN dentist 
							ON dentist.id = appointment.dentist_id 
							WHERE appointment.id = ?`

	row := s.db.QueryRow(queryGetById, id)

	appointment := domain.Appointment{}

	err := row.Scan(
		&appointment.Id,
		&appointment.Patient.Id,
		&appointment.Patient.Name,
		&appointment.Patient.LastName,
		&appointment.Patient.Rg,
		&appointment.Patient.RegisterDate,
		&appointment.Dentist.Id,
		&appointment.Dentist.Name,
		&appointment.Dentist.LastName,
		&appointment.Dentist.Enrollment,
		&appointment.Date,
		&appointment.Description,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return appointment, errors.New("appointment not found")
	}

	if err != nil {
		return appointment, err
	}

	return appointment, nil
}

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