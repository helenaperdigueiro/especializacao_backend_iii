package store

import "avaliacao-ii/internal/domain"

type StoreInterfacePatient interface {
	ReadById(id int) (domain.Patient, error)
	ReadByRg(rg string) (domain.Patient, error)
	ReadAll() ([]domain.Patient, error)
	Create(patient domain.Patient) (domain.Patient, error)
	Update(id int, patient domain.Patient) (domain.Patient, error)
	Patch(id int, patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type StoreInterfaceDentist interface {
	ReadById(id int) (domain.Dentist, error)
	ReadByEnrollment(enrollment string) (domain.Dentist, error)
	ReadAll() ([]domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Patch(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type StoreInterfaceAppointment interface {
	ReadById(id int) (domain.Appointment, error)
	ReadByRg(rg string) (domain.Appointment, error)
	// ReadAll() ([]domain.Appointment, error)
	// CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error)
	// CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollment string) (domain.Appointment, error)
	// Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	// Patch(id int, appointment domain.Appointment) (domain.Appointment, error)
	// Delete(id int) error
}