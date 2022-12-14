package appointment

import (
	"avaliacao-ii/internal/domain"
)

type Service interface {
	ReadById(id int) (domain.Appointment, error)
	ReadByRg(rg string) ([]domain.Appointment, error)
	CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error)
	// CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollment string) (domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Patch(id int, appointment domain.Appointment) (domain.Appointment, error)
	// Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ReadById(id int) (domain.Appointment, error) {
	appointment, err := s.r.ReadById(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) ReadByRg(rg string) ([]domain.Appointment, error) {
	appointments, err := s.r.ReadByRg(rg)
	if err != nil {
		return []domain.Appointment{}, err
	}
	return appointments, nil
}

func (s *service) CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {
	createdAppointment, err := s.r.CreateById(appointment, idPatient, idDentist)
	if err != nil {
		return domain.Appointment{}, err
	}
	return createdAppointment, nil
}

// func (s *service) CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollment string) (domain.Appointment, error) {
	
// }

func (s *service) Update(id int, appointment domain.Appointment) (domain.Appointment, error) {
	updatedAppointment, err := s.r.Update(id, appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return updatedAppointment, nil
}

func (s *service) Patch(id int, appointment domain.Appointment) (domain.Appointment, error) {
	updatedAppointment, err := s.r.Patch(id, appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return updatedAppointment, nil
}

// func (s *service) Delete(id int) error {
	
// }
