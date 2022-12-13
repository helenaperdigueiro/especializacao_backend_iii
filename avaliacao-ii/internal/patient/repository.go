package patient

import (
	"avaliacao-ii/internal/domain"
	"avaliacao-ii/pkg/store"
)

type Repository interface {
	ReadById(id int) (domain.Patient, error)
	ReadByRg(rg string) (domain.Patient, error)
	ReadAll() ([]domain.Patient, error)
	Create(patient domain.Patient) (domain.Patient, error)
	Update(id int, patient domain.Patient) (domain.Patient, error)
	Patch(id int, patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfacePatient
}

func NewRepository(storage store.StoreInterfacePatient) Repository {
	return &repository{storage}
}

func (r *repository) ReadById(id int) (domain.Patient, error) {
	patient, err := r.storage.ReadById(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (r *repository) ReadByRg(rg string) (domain.Patient, error) {
	patient, err := r.storage.ReadByRg(rg)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (r *repository) ReadAll() ([]domain.Patient, error) {
	patients, err := r.storage.ReadAll()
	if err != nil {
		return []domain.Patient{}, err
	}
	return patients, nil
}

func (r *repository) Create(patient domain.Patient) (domain.Patient, error) {
	createdPatient, err := r.storage.Create(patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return createdPatient, nil
}

func (r *repository) Update(id int, patient domain.Patient) (domain.Patient, error) {
	updatedPatient, err := r.storage.Update(id, patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return updatedPatient, nil
}

func (r *repository) Patch(id int, patient domain.Patient) (domain.Patient, error) {
	updatedPatient, err := r.storage.Patch(id, patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return updatedPatient, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
