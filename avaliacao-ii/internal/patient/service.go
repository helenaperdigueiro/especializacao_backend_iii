package patient

import (
	"errors"
	"avaliacao-ii/internal/domain"
)

type Service interface {
	ReadById(id int) (domain.Patient, error)
	Create(patient domain.Patient) (domain.Patient, error)
	Update(patient domain.Patient) (domain.Patient, error)
	Patch(patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ReadById(id int) (domain.Patient, error) {

}
func (s *service) Create(patient domain.Patient) (domain.Patient, error) {
	
}
func (s *service) Update(patient domain.Patient) (domain.Patient, error) {
	
}
func (s *service) Patch(patient domain.Patient) (domain.Patient, error) {
	
}
func (s *service) Delete(id int) error {
	
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
