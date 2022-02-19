package repositories

import (
	"context"
	"tiddly/src/configs"
	"tiddly/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ILinkRepository interface {
	GetLink(filter primitive.M) (*models.ShortLink, error)
	CreateLink(linkModel models.ShortLink) (*mongo.InsertOneResult, error)
}

type mongoClient struct {
	Client *mongo.Client
}

func LinkRepository(client *mongo.Client) ILinkRepository {
	return &mongoClient{Client: client}
}

func (c *mongoClient) GetLink(filter primitive.M) (*models.ShortLink, error) {
	shortLinkCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("shortlinks")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := &models.ShortLink{}

	if err := shortLinkCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *mongoClient) CreateLink(linkModel models.ShortLink) (*mongo.InsertOneResult, error) {
	shortLinkCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("shortlinks")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := shortLinkCollection.InsertOne(ctx, linkModel)
	if err != nil {
		return nil, err
	}

	return result, nil
}
