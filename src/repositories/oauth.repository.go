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

type IOAuthRepository interface {
	IsExistOAuth(id string) (*models.OauthClient, bool)
	CreateOAuth(oauthModel *models.OauthClient) (*mongo.InsertOneResult, error)
	UpdateOAuth(criteria primitive.M, update primitive.M) (*mongo.UpdateResult, error)
}

func OAuthRepository(client *mongo.Client) IOAuthRepository {
	return &mongoClient{client}
}

func (c *mongoClient) IsExistOAuth(id string) (*models.OauthClient, bool) {
	oauthCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("oauthclients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	oauth := &models.OauthClient{}
	if err := oauthCollection.FindOne(ctx, bson.M{"user_id": objId}).Decode(&oauth); err != nil {
		return nil, false
	}

	return oauth, true
}

func (c *mongoClient) CreateOAuth(oauthModel *models.OauthClient) (*mongo.InsertOneResult, error) {
	oauthCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("oauthclients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oauthModel.CreatedAt = time.Now()
	oauthModel.UpdatedAt = time.Now()

	result, err := oauthCollection.InsertOne(ctx, oauthModel)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *mongoClient) UpdateOAuth(criteria primitive.M, update primitive.M) (*mongo.UpdateResult, error) {
	oauthCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("oauthclients")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := oauthCollection.UpdateOne(ctx, criteria, update)
	if err != nil {
		return nil, err
	}

	return result, err
}
