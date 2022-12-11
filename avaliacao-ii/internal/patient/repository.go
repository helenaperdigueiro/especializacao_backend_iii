package patient

import (
	"errors"
	"avaliacao-ii/internal/domain"
	"avaliacao-ii/pkg/store"
)

type Repository interface {
	ReadById(id int) (domain.Patient, error)
	Create(patient domain.Patient) (domain.Patient, error)
	Update(patient domain.Patient) (domain.Patient, error)
	Patch(patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfacePatient
}

func NewRepository(storage store.StoreInterfacePatient) Repository {
	return &repository{storage}
}

func (r *repository) ReadById(id int) (domain.Patient, error) {

}
func (r *repository) Create(patient domain.Patient) (domain.Patient, error) {
	
}
func (r *repository) Update(patient domain.Patient) (domain.Patient, error) {
	
}
func (r *repository) Patch(patient domain.Patient) (domain.Patient, error) {
	
}
func (r *repository) Delete(id int) error {
	
}

func (r *repository) GetByID(id int) (domain.Product, error) {
	product, err := r.storage.Read(id)
	if err != nil {
		return domain.Product{}, errors.New("product not found")
	}
	return product, nil

}

func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.storage.Exists(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	err := r.storage.Create(p)
	if err != nil {
		return domain.Product{}, errors.New("error creating product")
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, p domain.Product) (domain.Product, error) {
	if !r.storage.Exists(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	err := r.storage.Update(p)
	if err != nil {
		return domain.Product{}, errors.New("error updating product")
	}
	return p, nil
}
