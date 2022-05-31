package product

import (
	"errors"

	"github.com/danisbagus/golang-hexagon-mongo/core/model"
	port "github.com/danisbagus/golang-hexagon-mongo/core/port/product"
	portTransactor "github.com/danisbagus/golang-hexagon-mongo/core/port/transactor"
)

type Service struct {
	repo       port.Repository
	transactor portTransactor.Transactor
}

func New(repo port.Repository) port.Service {
	return &Service{
		repo: repo,
	}
}

func (s Service) Insert(form *model.Product) error {
	return s.repo.Insert(form)
}

func (s Service) View(ID uint64) (*model.Product, error) {
	product, err := s.repo.FindOneByID(ID)
	if err == nil {
		return nil, err
	}

	if product.ID == 0 {
		return nil, errors.New("product not found")
	}

	return product, nil
}
