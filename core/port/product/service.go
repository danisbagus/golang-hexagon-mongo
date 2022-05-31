package product

import "github.com/danisbagus/golang-hexagon-mongo/core/model"

type Service interface {
	// Insert
	Insert(form *model.Product) error

	// View
	View(ID uint64) (*model.Product, error)
}
