package Repository

import (
	"context"
	"fmt"

	"github.com/danisbagus/golang-hexagon-mongo/core/model"
	port "github.com/danisbagus/golang-hexagon-mongo/core/port/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = "products"

type Repository struct {
	coll *mongo.Collection
}

type Product struct {
	ID         string `bson:"_id,omitempty"`
	Name       string `bson:"name"`
	CategoryID uint64 `bson:"category_id"`
	Price      uint64 `bson:"price"`
}

func New(db *mongo.Database) port.Repository {
	return &Repository{
		coll: db.Collection(collectionName),
	}
}

func (r Repository) Insert(inData *model.Product) error {
	product := newProduct(inData)

	res, err := r.coll.InsertOne(context.Background(), product)
	if err != nil {
		return fmt.Errorf("failed insert product: %v", err)
	}

	if res.InsertedID == "" {
		return fmt.Errorf("failed insert product: no data was inserted")
	}

	return nil
}

func (r Repository) FindAll() ([]model.Product, error) {
	products := make([]Product, 0)
	productsOut := make([]model.Product, 0)
	filter := bson.M{}

	cursor, err := r.coll.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed get list product: %v", err)
	}

	err = cursor.All(context.Background(), &products)
	if err != nil {
		return nil, fmt.Errorf("failed read all cursor: %v", err)
	}

	for _, product := range products {
		var productOut model.Product
		productOut.ID = product.ID
		productOut.Name = product.Name
		productOut.CategoryID = product.CategoryID
		productOut.Price = product.Price

		productsOut = append(productsOut, productOut)
	}

	return productsOut, nil
}

func (r Repository) FindOneByID(ID string) (*model.Product, error) {
	oid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, fmt.Errorf("failed convert object id: %v", err)
	}

	product := new(Product)
	filter := bson.M{"_id": oid}
	res := r.coll.FindOne(context.Background(), filter)
	if err := res.Decode(product); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("failed get product: data not found")

		}
		return nil, fmt.Errorf("failed get product: %v", err)
	}

	productOut := newProductOut(product)
	return productOut, nil
}

func newProduct(inData *model.Product) *Product {
	product := new(Product)
	product.ID = inData.ID
	product.Name = inData.Name
	product.CategoryID = inData.CategoryID
	product.Price = inData.Price
	return product
}

func newProductOut(product *Product) *model.Product {
	outData := new(model.Product)
	outData.ID = product.ID
	outData.Name = product.Name
	outData.CategoryID = product.CategoryID
	outData.Price = product.Price
	return outData
}
