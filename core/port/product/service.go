package product

import "github.com/danisbagus/golang-hexagon-mongo/core/model"

type Service interface {
	// Insert
	Insert(product *model.Product) error

	// List
	List() ([]model.Product, error)

	// View
	View(ID string) (*model.Product, error)

	// Update
	Update(product *model.Product) error
}
