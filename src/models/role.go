package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}
