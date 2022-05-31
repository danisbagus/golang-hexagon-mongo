package product

import "github.com/danisbagus/golang-hexagon-mongo/core/model"

type Repository interface {
	// Insert
	Insert(product *model.Product) error

	// Find one
	FindOneByID(ID uint64) (*model.Product, error)
}
