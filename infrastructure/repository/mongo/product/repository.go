package Repository

import (
	"context"
	"fmt"
	"time"

	"github.com/danisbagus/golang-hexagon-mongo/core/model"
	port "github.com/danisbagus/golang-hexagon-mongo/core/port/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = "products"

type Repository struct {
	coll   *mongo.Collection
	client *mongo.Client
}

type Product struct {
	ID          primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	CategoryIDs []uint64           `json:"category_ids" bson:"category_ids"`
	Price       uint64             `json:"price" bson:"price"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

func New(db *mongo.Database, client *mongo.Client) port.Repository {
	return &Repository{
		coll:   db.Collection(collectionName),
		client: client,
	}
}

func (r Repository) Insert(ctx context.Context, inData *model.Product) error {
	product := newProduct(inData)
	res, err := r.coll.InsertOne(ctx, product)
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
		productOut.ID = product.ID.Hex()
		productOut.Name = product.Name
		productOut.CategoryIDs = product.CategoryIDs
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

func (r Repository) Update(ctx context.Context, inData *model.Product) error {
	oid, err := primitive.ObjectIDFromHex(inData.ID)
	if err != nil {
		return fmt.Errorf("failed convert object id: %v", err)
	}

	product := newProduct(inData)
	filter := bson.M{"_id": oid}

	result, err := r.coll.UpdateOne(ctx, filter, bson.M{"$set": product})
	if err != nil {
		return fmt.Errorf("failed update product: %v", err)
	}

	if result.ModifiedCount < 1 {
		return fmt.Errorf("failed update product: no data was updated")
	}

	return nil
}

func (r Repository) Delete(ctx context.Context, ID string) error {
	oid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return fmt.Errorf("failed convert object id: %v", err)
	}

	res, err := r.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return fmt.Errorf("failed delete product: %v", err)
	}

	if res.DeletedCount < 1 {
		return fmt.Errorf("failed delete product: no data was deleted")
	}

	return nil

}

func newProduct(inData *model.Product) *Product {
	product := new(Product)
	oid, _ := primitive.ObjectIDFromHex(inData.ID)
	product.ID = oid
	product.Name = inData.Name
	product.CategoryIDs = inData.CategoryIDs
	product.Price = inData.Price
	product.CreatedAt = inData.CreatedAt
	return product
}

func newProductOut(product *Product) *model.Product {
	outData := new(model.Product)
	outData.ID = product.ID.Hex()
	outData.Name = product.Name
	outData.CategoryIDs = product.CategoryIDs
	outData.Price = product.Price
	outData.CreatedAt = product.CreatedAt
	return outData
}
