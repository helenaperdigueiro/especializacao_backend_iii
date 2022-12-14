package store

import (
	"avaliacao-ii/internal/domain"
	"database/sql"
	"errors"
	"log"
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

func (s *sqlStoreAppointment) ReadByRg(rg string) ([]domain.Appointment, error) {
	queryGetByRg := `SELECT appointment.id, patient.id, patient.name, patient.last_name, patient.rg, patient.register_date, 
							dentist.id, dentist.name, dentist.last_name, dentist.enrollment, 
							appointment.date, appointment.description 
							FROM appointment 
							INNER JOIN patient 
							ON patient.id = appointment.patient_id 
							INNER JOIN dentist 
							ON dentist.id = appointment.dentist_id 
							WHERE patient.rg = ?`

	var appointments []domain.Appointment
	rows, err := s.db.Query(queryGetByRg, rg)
	if err != nil {
		return []domain.Appointment{}, err
	}

	defer rows.Close()

	for rows.Next() {
		appointment := domain.Appointment{}

		if err := rows.Scan(
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
		); err != nil {
			return appointments, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (s *sqlStoreAppointment) CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {
	queryInsert := "INSERT INTO appointment (patient_id, dentist_id, date, description) VALUES (?, ?, ?, ?)"

	stmt, err := s.db.Prepare(queryInsert)

	if err != nil {
		return domain.Appointment{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		idPatient,
		idDentist,
		appointment.Date,
		appointment.Description)
	if err != nil {
		return domain.Appointment{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return domain.Appointment{}, errors.New("fail to save appointment")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return domain.Appointment{}, err
	}

	appointment.Id = int(lastId)

	appointment, err = s.ReadById(appointment.Id)
	if err != nil {
		return domain.Appointment{}, err
	}

	return appointment, nil
}

// func (s *sqlStoreAppointment) CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollment string) (domain.Appointment, error) {

// }

func (s *sqlStoreAppointment) Update(id int, appointment domain.Appointment) (domain.Appointment, error) {
	queryUpdate  := "UPDATE appointment SET patient_id = ?, dentist_id = ?, date = ?, description = ? WHERE id = ?"

	persistedAppointment, err := s.ReadById(id)
	if err != nil {
		return domain.Appointment{}, errors.New("appointment not found")
	}

	persistedAppointment.Patient.Id = appointment.Patient.Id
	persistedAppointment.Dentist.Id = appointment.Dentist.Id
	persistedAppointment.Date = appointment.Date
	persistedAppointment.Description = appointment.Description

	result, err := s.db.Exec(
		queryUpdate,
		persistedAppointment.Patient.Id,
		persistedAppointment.Dentist.Id,
		persistedAppointment.Date,
		persistedAppointment.Description,
		id,
	)
	if err != nil {
		return domain.Appointment{}, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Appointment{}, err
	}
	log.Println(affectedRows)

	appointment, err = s.ReadById(id)
	if err != nil {
		return domain.Appointment{}, err
	}

	return appointment, nil
}

func (s *sqlStoreAppointment) Patch(id int, appointment domain.Appointment) (domain.Appointment, error) {
	queryUpdate  := "UPDATE appointment SET patient_id = ?, dentist_id = ?, date = ?, description = ? WHERE id = ?"

	persistedAppointment, err := s.ReadById(id)
	if err != nil {
		return domain.Appointment{}, errors.New("appointment not found")
	}

	if appointment.Patient.Id != 0 {
		persistedAppointment.Patient.Id = appointment.Patient.Id
	}
	if appointment.Dentist.Id != 0 {
		persistedAppointment.Dentist.Id = appointment.Dentist.Id
	}
	if appointment.Date != "" {
		persistedAppointment.Date = appointment.Date
	}
	if appointment.Description != "" {
		persistedAppointment.Description = appointment.Description
	}

	result, err := s.db.Exec(
		queryUpdate,
		persistedAppointment.Patient.Id,
		persistedAppointment.Dentist.Id,
		persistedAppointment.Date,
		persistedAppointment.Description,
		id,
	)
	if err != nil {
		return domain.Appointment{}, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Appointment{}, err
	}
	log.Println(affectedRows)

	appointment, err = s.ReadById(id)
	if err != nil {
		return domain.Appointment{}, err
	}

	return appointment, nil
}

// func (s *sqlStoreAppointment) Delete(id int) error {

// }
