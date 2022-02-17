package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BillPaymentItem struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	AppPackageId primitive.ObjectID `bson:"app_package_id"`
	Quantity     uint               `bson:"quantity,omitempty"`
	Price        float32            `bson:"price,omitempty"`
	CreatedAt    time.Time          `bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at,omitempty"`
}
