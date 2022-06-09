package product

import (
	"context"
	"fmt"
	"time"

	"github.com/danisbagus/golang-hexagon-mongo/core/model"
	port "github.com/danisbagus/golang-hexagon-mongo/core/port/product"
	portTransactor "github.com/danisbagus/golang-hexagon-mongo/core/port/transactor"
)

type Service struct {
	repo       port.Repository
	transactor portTransactor.Transactor
}

func New(repo port.Repository, transactor portTransactor.Transactor) port.Service {
	return &Service{
		repo:       repo,
		transactor: transactor,
	}
}

func (s Service) Insert(form *model.Product) error {
	return s.transactor.WithinTransaction(func(txCtx context.Context) error {
		timeNow := time.Now()
		form.CreatedAt = timeNow
		err := s.repo.Insert(txCtx, form)
		if err != nil {
			return err
		}
		return nil
	})

}

func (s Service) List() ([]model.Product, error) {
	products, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s Service) View(ID string) (*model.Product, error) {
	product, err := s.repo.FindOneByID(ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s Service) Update(form *model.Product) error {
	return s.transactor.WithinTransaction(func(txCtx context.Context) error {
		product, err := s.repo.FindOneByID(form.ID)
		if err != nil {
			return err
		}
		if product.ID == "" {
			return fmt.Errorf("product not found")
		}

		err = s.repo.Update(txCtx, form)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s Service) Delete(ID string) error {
	return s.transactor.WithinTransaction(func(txCtx context.Context) error {
		product, err := s.repo.FindOneByID(ID)
		if err != nil {
			return err
		}
		if product.ID == "" {
			return fmt.Errorf("product not found")
		}

		err = s.repo.Delete(txCtx, ID)
		if err != nil {
			return err
		}
		return nil
	})
}
