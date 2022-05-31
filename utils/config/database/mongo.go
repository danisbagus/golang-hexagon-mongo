package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoClient   *mongo.Client
	MongoDatabase *mongo.Database
)

type MongoConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	Url      string
}

func LoadMongoConfig() *MongoConfig {
	userName := "root"
	password := "pwd123"
	host := "localhost"
	port := 8800
	dbName := "simple-prodoct-management"

	url := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		userName,
		password,
		host,
		port,
	)

	mongoConfig := &MongoConfig{
		User:     userName,
		Password: password,
		Host:     host,
		Port:     port,
		DBName:   dbName,
		Url:      url,
	}

	return mongoConfig

}

func OpenMongoConnection() {

	if MongoDatabase != nil {
		return
	}

	mongoConfig := LoadMongoConfig()

	client, err := NewClient(mongoConfig)
	if err != nil {
		log.Fatalf("Failed to connect MongoDB %v", err)
	}

	MongoClient = client
	MongoDatabase = client.Database(mongoConfig.DBName)
}

func NewClient(mongoConfig *MongoConfig) (*mongo.Client, error) {
	if MongoClient != nil {
		return MongoClient, nil
	}

	url := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		mongoConfig.User,
		mongoConfig.Password,
		mongoConfig.Host,
		mongoConfig.Port,
	)

	mongoOptions := options.Client().ApplyURI(url)

	client, err := mongo.NewClient(mongoOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		client.Disconnect(context.TODO())
		return nil, err
	}

	return client, nil

}
