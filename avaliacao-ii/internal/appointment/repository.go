package appointment

import (
	"avaliacao-ii/internal/domain"
	"avaliacao-ii/pkg/store"
)

type Repository interface {
	ReadById(id int) (domain.Appointment, error)
	ReadByRg(rg string) ([]domain.Appointment, error)
	CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error)
	// CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollmentDentist string) (domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Patch(id int, appointment domain.Appointment) (domain.Appointment, error)
	// Delete(id int) error
}

type repository struct {
	storage store.StoreInterfaceAppointment
}

func NewRepository(storage store.StoreInterfaceAppointment) Repository {
	return &repository{storage}
}

func (r *repository) ReadById(id int) (domain.Appointment, error) {
	appointment, err := r.storage.ReadById(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (r *repository) ReadByRg(rg string) ([]domain.Appointment, error) {
	appointments, err := r.storage.ReadByRg(rg)
	if err != nil {
		return []domain.Appointment{}, err
	}
	return appointments, nil
}

func (r *repository) CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {
	createdAppointment, err := r.storage.CreateById(appointment, idPatient, idDentist)
	if err != nil {
		return domain.Appointment{}, err
	}
	return createdAppointment, nil
}

// func (r *repository) CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollmentDentist string) (domain.Appointment, error) {
	
// }

func (r *repository) Update(id int, appointment domain.Appointment) (domain.Appointment, error) {
	updatedAppointment, err := r.storage.Update(id, appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return updatedAppointment, nil
}

func (r *repository) Patch(id int, appointment domain.Appointment) (domain.Appointment, error) {
	updatedAppointment, err := r.storage.Patch(id, appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return updatedAppointment, nil
}

// func (r *repository) Delete(id int) error {
	
// }
