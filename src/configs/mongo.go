package configs

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConectToMongo() (*mongo.Client , context.Context,error) {
	fmt.Print(loadEnv("MONGO_DB"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(loadEnv("MONGO_DB")))
	if err != nil {
		panic(err)
	}

	fmt.Printf("DATABASE CONNECTed")

	return client,ctx,nil
}
