package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type App struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	AppName   string             `bson:"app_name,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}
