package repositories

import "go.mongodb.org/mongo-driver/mongo"

type IAppRepository interface {
	ListAppPackage()
	GetAppPackage()
	CreateAppPackage()
	UpdateAppPackage()
	DeleteAppPackage()
}

func AppRepository(client *mongo.Client) IAppRepository {
	return &mongoClient{client}
}

func (c *mongoClient) ListAppPackage() {

}

func (c *mongoClient) GetAppPackage() {

}

func (c *mongoClient) CreateAppPackage() {

}

func (c *mongoClient) UpdateAppPackage() {

}

func (c *mongoClient) DeleteAppPackage() {

}
