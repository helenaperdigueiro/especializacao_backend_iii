package appointment

import (
	"errors"
	"avaliacao-ii/internal/domain"
	"avaliacao-ii/pkg/store"
)

type Repository interface {
	ReadById(id int) (domain.Appointment, error)
	ReadByRg(rg string) (domain.Appointment, error)
	CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error)
	CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollment string) (domain.Appointment, error)
	Update(appointment domain.Appointment) (domain.Appointment, error)
	Patch(appointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfaceAppointment
}

func NewRepository(storage store.StoreInterfaceAppointment) Repository {
	return &repository{storage}
}

func (r *repository) ReadById(id int) (domain.Appointment, error) {

}
func (r *repository) ReadByRg(rg string) (domain.Appointment, error) {
	
}
func (r *repository) CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {
	
}
func (r *repository) CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollment string) (domain.Appointment, error) {
	
}
func (r *repository) Update(appointment domain.Appointment) (domain.Appointment, error) {
	
}
func (r *repository) Patch(appointment domain.Appointment) (domain.Appointment, error) {
	
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
