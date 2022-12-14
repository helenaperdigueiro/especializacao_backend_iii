package product

import (
	"errors"

	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
)

const (
	impuesto21Percent = 1.21
	impuesto17Percent = 1.17
	impuesto15Percent = 1.15
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) ([]domain.Product, error)
	Create(p domain.Product) (domain.Product, error)
	Update(id int, p domain.Product) (domain.Product, error)
	Delete(id int) error
	ConsumerPrice(ids []int) (domain.ConsumerPriceResponse, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetAll devuelve todos los productos
func (s *service) GetAll() ([]domain.Product, error) {
	l := s.r.GetAll()
	return l, nil
}

// GetByID busca un producto por su id
func (s *service) GetByID(id int) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func countIds(ids []int) (result map[int]int) {
	result = make(map[int]int)
	for _, id := range ids {
		if v, found := result[id]; found {
			result[id] = v + 1
		} else {
			result[id] = 1
		}
	}
	return
}

// SearchPriceGt busca productos por precio mayor que el precio dado
func (s *service) SearchPriceGt(price float64) ([]domain.Product, error) {
	l := s.r.SearchPriceGt(price)
	if len(l) == 0 {
		return []domain.Product{}, errors.New("no products found")
	}
	return l, nil
}

// Create agrega un nuevo producto
func (s *service) Create(p domain.Product) (domain.Product, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

// Delete elimina un producto
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update actualiza un producto
func (s *service) Update(id int, u domain.Product) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	if u.Name != "" {
		p.Name = u.Name
	}
	if u.CodeValue != "" {
		p.CodeValue = u.CodeValue
	}
	if u.Expiration != "" {
		p.Expiration = u.Expiration
	}
	if u.Quantity > 0 {
		p.Quantity = u.Quantity
	}
	if u.Price > 0 {
		p.Price = u.Price
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *service) ConsumerPrice(ids []int) (domain.ConsumerPriceResponse, error) {
	var products []domain.Product
	var total float64
	for _, id := range ids {
		product, err := s.r.GetByID(id)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
		total += product.Price
	}

	totalWithTax := total * impuesto21Percent

	consumerPriceResponse := domain.ConsumerPriceResponse {
		Products : products,
		TotalPrice : totalWithTax,
	}

	return consumerPriceResponse, nil
}
