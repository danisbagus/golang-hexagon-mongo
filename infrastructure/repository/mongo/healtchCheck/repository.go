package Repository

import (
	"context"

	port "github.com/danisbagus/golang-hexagon-mongo/core/port/healthCheck"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	client *mongo.Client
}

func New(client *mongo.Client) port.Repository {
	return &Repository{
		client: client,
	}
}

func (r Repository) Ping() error {
	return r.client.Ping(context.Background(), nil)
}
