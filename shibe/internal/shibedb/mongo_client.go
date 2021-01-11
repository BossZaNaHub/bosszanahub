package shibedb

import "go.mongodb.org/mongo-driver/mongo"

type defaultClient struct {
	db *mongo.Database
}

func NewClient(db *mongo.Database) Client {
	return &defaultClient{
		db: db,
	}
}