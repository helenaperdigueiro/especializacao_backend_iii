package dentist

import (
	"errors"
	"avaliacao-ii/internal/domain"
)

type Service interface {
	ReadById(id int) (domain.Dentist, error)
	ReadByEnrollment(enrollment string) (domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Patch(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ReadById(id int) (domain.Dentist, error) {
	dentist, err := s.r.ReadById(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *service) ReadByEnrollment(enrollment string) (domain.Dentist, error) {
	dentist, err := s.r.ReadByEnrollment(enrollment)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *service) Create(dentist domain.Dentist) (domain.Dentist, error) {
	persistedDentist, err := s.ReadByEnrollment(dentist.Enrollment)
	if dentist.Enrollment == persistedDentist.Enrollment {
		return domain.Dentist{}, errors.New("enrollment already exists")
	}

	createdDentist, err := s.r.Create(dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return createdDentist, nil
}

func (s *service) Update(id int, dentist domain.Dentist) (domain.Dentist, error) {
	persistedDentist, err := s.ReadByEnrollment(dentist.Enrollment)
	if dentist.Enrollment == persistedDentist.Enrollment {
		return domain.Dentist{}, errors.New("enrollment already exists")
	}

	updatedDentist, err := s.r.Update(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return updatedDentist, nil
}

func (s *service) Patch(id int, dentist domain.Dentist) (domain.Dentist, error) {
	persistedDentist, err := s.ReadByEnrollment(dentist.Enrollment)
	if dentist.Enrollment == persistedDentist.Enrollment {
		return domain.Dentist{}, errors.New("enrollment already exists")
	}

	updatedDentist, err := s.r.Patch(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return updatedDentist, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
