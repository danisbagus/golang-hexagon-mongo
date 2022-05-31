package product

import "github.com/danisbagus/golang-hexagon-mongo/core/model"

type Service interface {
	// Insert
	Insert(form *model.Product) error

	// View
	View(ID string) (*model.Product, error)
}
