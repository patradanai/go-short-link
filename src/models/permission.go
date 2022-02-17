package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Permission struct {
	Id        primitive.ObjectID `bson:"_id"`
	MenuId    primitive.ObjectID `bson:"menu_id"`
	Name      string             `bson:"name"`
	Action    string             `bson:"action"`
	Status    bool               `bson:"status"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
