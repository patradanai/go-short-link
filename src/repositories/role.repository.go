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

type IRoleRepository interface {
	ListRole() ([]models.Role, error)
	CreateRole(roleModel *models.Role) (*mongo.InsertOneResult, error)
	UpdateRole(id string, roleModel *models.Role) (*mongo.UpdateResult, error)
	DeleteRole(id string) (*mongo.DeleteResult, error)
}

func RoleRepository(client *mongo.Client) IRoleRepository {
	return &mongoClient{Client: client}
}

func (c *mongoClient) ListRole() ([]models.Role, error) {
	roleCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("roles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := roleCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	roles := make([]models.Role, 0)
	for cur.Next(ctx) {
		role := models.Role{}
		if err := cur.Decode(&role); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (c *mongoClient) CreateRole(roleModel *models.Role) (*mongo.InsertOneResult, error) {
	roleCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("roles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	role := roleModel
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()
	result, err := roleCollection.InsertOne(ctx, role)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *mongoClient) UpdateRole(id string, roleModel *models.Role) (*mongo.UpdateResult, error) {
	roleCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("roles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	roleModel.CreatedAt = time.Now()
	roleModel.UpdatedAt = time.Now()

	update := bson.M{"$set": roleModel}
	result, err := roleCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *mongoClient) DeleteRole(id string) (*mongo.DeleteResult, error) {
	roleCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("roles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, err := roleCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return nil, err
	}

	return result, nil
}
