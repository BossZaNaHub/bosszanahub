package bzmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *defaultClient) ListAllCollections() (int, []string) {

	names, _ := c.DB.ListCollectionNames(context.TODO(), bson.D{})
	count := len(names)

	return count, names
}

func (c *defaultClient) SelectCollection(collection string, filter interface{}) (interface{}, error){
	data, err := c.DB.Collection(collection).Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	val := data.Decode(&filter)
	return val, nil
}