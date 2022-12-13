package appointment

import (
	"avaliacao-ii/internal/domain"
	"avaliacao-ii/pkg/store"
)

type Repository interface {
	ReadById(id int) (domain.Appointment, error)
	// CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error)
	// CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollmentDentist string) (domain.Appointment, error)
	// Update(appointment domain.Appointment) (domain.Appointment, error)
	// Patch(appointment domain.Appointment) (domain.Appointment, error)
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

// func (r *repository) CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {
	
// }

// func (r *repository) CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollmentDentist string) (domain.Appointment, error) {
	
// }

// func (r *repository) Update(appointment domain.Appointment) (domain.Appointment, error) {
	
// }

// func (r *repository) Patch(appointment domain.Appointment) (domain.Appointment, error) {
	
// }

// func (r *repository) Delete(id int) error {
	
// }
