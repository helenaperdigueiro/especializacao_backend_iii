package appointment

import (
	"avaliacao-ii/internal/domain"
)

type Service interface {
	ReadById(id int) (domain.Appointment, error)
	ReadByRg(rg string) (domain.Appointment, error)
	// CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error)
	// CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollment string) (domain.Appointment, error)
	// Update(appointment domain.Appointment) (domain.Appointment, error)
	// Patch(appointment domain.Appointment) (domain.Appointment, error)
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

func (s *service) ReadByRg(rg string) (domain.Appointment, error) {
	appointment, err := s.r.ReadByRg(rg)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

// func (s *service) CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {
	
// }

// func (s *service) CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollment string) (domain.Appointment, error) {
	
// }

// func (s *service) Update(appointment domain.Appointment) (domain.Appointment, error) {
	
// }

// func (s *service) Patch(appointment domain.Appointment) (domain.Appointment, error) {
	
// }

// func (s *service) Delete(id int) error {
	
// }
