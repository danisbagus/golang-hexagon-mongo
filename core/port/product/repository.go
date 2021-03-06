package product

import (
	"context"

	"github.com/danisbagus/golang-hexagon-mongo/core/model"
)

type Repository interface {
	// Insert
	Insert(ctx context.Context, product *model.Product) error

	// Find all
	FindAll() ([]model.Product, error)

	// Find one
	FindOneByID(ID string) (*model.Product, error)

	// Update
	Update(ctx context.Context, product *model.Product) error

	// Delete
	Delete(ctx context.Context, ID string) error
}
