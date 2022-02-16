package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type imageInfo struct {
	Src   string `bson:"src,omitempty"`
	Title string `bson:"title,omitempty"`
}
type corporateInfo struct {
	NameTh string    `bson:"name_th,omitempty"`
	NameEn string    `bson:"name_en,omitempty"`
	Phone  string    `bson:"phone,omitempty"`
	Image  imageInfo `bson:"image,omitempty"`
}
type Corporate struct {
	Id            primitive.ObjectID   `bson:"_id,omitempty"`
	AppId         primitive.ObjectID   `bson:"app_id"`
	CorporateInfo corporateInfo        `bson:"corporate_info"`
	BillPayments  []primitive.ObjectID `bson:"bill_payments,omitempty"`
	Packages      []string             `bson:"packages,omitempty"`
	CreatedAt     time.Time            `bson:"created_at,omitempty"`
	UpdatedAt     time.Time            `bson:"updated_at,omitempty"`
}
