package dentist

import (
	"avaliacao-ii/internal/domain"
	"avaliacao-ii/pkg/store"
)

type Repository interface {
	ReadById(id int) (domain.Dentist, error)
	ReadByEnrollment(enrollment string) (domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Patch(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfaceDentist
}

func NewRepository(storage store.StoreInterfaceDentist) Repository {
	return &repository{storage}
}

func (r *repository) ReadById(id int) (domain.Dentist, error) {
	dentist, err := r.storage.ReadById(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (r *repository) ReadByEnrollment(enrollment string) (domain.Dentist, error) {
	dentist, err := r.storage.ReadByEnrollment(enrollment)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (r *repository) Create(dentist domain.Dentist) (domain.Dentist, error) {
	createdDentist, err := r.storage.Create(dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return createdDentist, nil
}

func (r *repository) Update(id int, dentist domain.Dentist) (domain.Dentist, error) {
	updatedDentist, err := r.storage.Update(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return updatedDentist, nil
}

func (r *repository) Patch(id int, dentist domain.Dentist) (domain.Dentist, error) {
	updatedDentist, err := r.storage.Patch(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return updatedDentist, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
