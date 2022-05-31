package product

import (
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

func (s Service) View(ID string) (*model.Product, error) {
	product, err := s.repo.FindOneByID(ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}
