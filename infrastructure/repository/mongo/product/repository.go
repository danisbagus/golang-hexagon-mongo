package Repository

import (
	"context"
	"fmt"

	"github.com/danisbagus/golang-hexagon-mongo/core/model"
	port "github.com/danisbagus/golang-hexagon-mongo/core/port/product"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Database
}

type Product struct {
	ID         uint64 `bson:"_id,omitempty"`
	Name       string `bson:"name"`
	CategoryID uint64 `bson:"category_id"`
	Price      uint64 `bson:"price"`
}

func New(db *mongo.Database) port.Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Insert(inData *model.Product) error {
	product := toProdoct(inData)

	res, err := r.db.Collection("products").InsertOne(context.Background(), product)
	if err != nil {
		return fmt.Errorf("failed insert product: %v", err.Error())
	}

	if res.InsertedID == "" {
		return fmt.Errorf("failed insert product: no data was inserted")
	}

	return nil
}

func (r Repository) FindOneByID(ID uint64) (*model.Product, error) {
	return nil, nil
}

func toProdoct(inData *model.Product) *Product {
	product := new(Product)
	product.ID = inData.ID
	product.Name = inData.Name
	product.CategoryID = inData.CategoryID
	product.Price = inData.Price
	return product
}

func toProductOut(product *Product) *model.Product {
	outData := new(model.Product)
	outData.ID = product.ID
	outData.Name = product.Name
	outData.CategoryID = product.CategoryID
	outData.Price = product.Price
	return outData
}
