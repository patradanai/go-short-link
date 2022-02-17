package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type packageInfo struct {
	PackageSize string    `bson:"package_size,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty"`
}
type AppPackage struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	AppId        primitive.ObjectID `bson:"app_id,omitempty"`
	PacakageName string             `bson:"package_name,omitempty"`
	PackageInfo  packageInfo        `bson:"package_info,omitempty"`
	CreatedAt    time.Time          `bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at,omitempty"`
}
