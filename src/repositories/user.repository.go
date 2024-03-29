package repositories

import (
	"context"
	"tiddly/src/configs"
	"tiddly/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	UpdateUserInfo(id string, filter primitive.M) (*mongo.UpdateResult, error)
	GetUser(id string) (*models.User, error)
}

func UserRepository(client *mongo.Client) IUserRepository {
	return &mongoClient{client}
}

func (c *mongoClient) UpdateUserInfo(id string, filter primitive.M) (*mongo.UpdateResult, error) {
	userCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result, err := userCollection.UpdateByID(ctx, objId, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *mongoClient) GetUser(id string) (*models.User, error) {
	userCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	userModel := &models.User{}

	if err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&userModel); err != nil {
		return nil, err
	}

	return userModel, nil
}
