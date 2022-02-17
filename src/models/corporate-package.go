package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CorporatePackage struct {
	Id               primitive.ObjectID `bson:"_id,omitempty"`
	AppPackageId     primitive.ObjectID `bson:"app_package_id"`
	CorporateId      primitive.ObjectID `bson:"corporate_id"`
	Status           bool               `bson:"status,omitempty"`
	StatusReason     string             `bson:"status_reason,omitempty"`
	PackageSize      uint               `bson:"package_size,omitempty"`
	PackageSizeLimit uint               `bson:"package_size_limit,omitempty"`
	ExpiredAt        time.Time          `bson:"expired_at"`
	CreatedAt        time.Time          `bson:"created_at,omitempty"`
	UpdatedAt        time.Time          `bson:"updated_at,omitempty"`
}
