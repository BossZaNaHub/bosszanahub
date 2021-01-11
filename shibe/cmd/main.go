package main

import (
	"context"
	"fmt"
	"github.com/bosszanahub/shibe/internal/shibedb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(""))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)

	mongo := shibedb.NewClient(client.Database(""))
	fmt.Println(mongo)
}