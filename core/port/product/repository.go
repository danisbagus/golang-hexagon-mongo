package product

import "github.com/danisbagus/golang-hexagon-mongo/core/model"

type Repository interface {
	// Insert
	Insert(product *model.Product) error

	// Find all
	FindAll() ([]model.Product, error)

	// Find one
	FindOneByID(ID string) (*model.Product, error)
}
