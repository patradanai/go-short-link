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

type IAppRepository interface {
	ListAppPackage() ([]models.AppPackage, error)
	GetAppPackage(id string) (*models.AppPackage, error)
	CreateAppPackage(packageModel *models.AppPackage) (*mongo.InsertOneResult, error)
	UpdateAppPackage()
	DeleteAppPackage(id string) (*mongo.DeleteResult, error)
}

func AppRepository(client *mongo.Client) IAppRepository {
	return &mongoClient{client}
}

func (c *mongoClient) ListAppPackage() ([]models.AppPackage, error) {
	appPackageCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("apppackages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := appPackageCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	packageArray := make([]models.AppPackage, 0)

	for cur.Next(ctx) {
		packageOne := models.AppPackage{}

		if err := cur.Decode(&packageOne); err != nil {
			return nil, err
		}

		packageArray = append(packageArray, packageOne)
	}

	return packageArray, nil
}

func (c *mongoClient) GetAppPackage(id string) (*models.AppPackage, error) {
	appPackageCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("apppackages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	appPackage := models.AppPackage{}
	if err := appPackageCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&appPackage); err != nil {
		return nil, err
	}

	return &appPackage, nil
}

func (c *mongoClient) CreateAppPackage(packageModel *models.AppPackage) (*mongo.InsertOneResult, error) {
	appPackageCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("apppackages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	packageModel.CreatedAt = time.Now()
	packageModel.UpdatedAt = time.Now()

	result, err := appPackageCollection.InsertOne(ctx, packageModel)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *mongoClient) UpdateAppPackage() {

}

func (c *mongoClient) DeleteAppPackage(id string) (*mongo.DeleteResult, error) {
	appPackageCollection := c.Client.Database(configs.LoadEnv("MONGO_DB_NAME")).Collection("apppackages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, err := appPackageCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return nil, err
	}

	return result, nil
}
