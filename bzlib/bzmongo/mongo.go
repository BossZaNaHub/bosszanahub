package bzmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type defaultClient struct {
	Uri string
	Database string
	Timeout time.Duration
	Client *mongo.Client
	DB *mongo.Database
	MongoClient MongoClient
}

func NewMongo(uri, databaseName string, timeout time.Duration) MongoClient {
	return &defaultClient{
		Uri:      uri,
		Database: databaseName,
		Timeout:  timeout,
	}
}

func (c *defaultClient) OpenMongoConnection() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(c.Uri))

	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), c.Timeout * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	c.DB = client.Database(c.Database)
	c.Client = client

	return nil
}

func (c *defaultClient) CloseConnection() {
	c.Client.Disconnect(context.Background())
}