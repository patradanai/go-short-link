package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OauthClient struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	UserId       primitive.ObjectID `bson:"user_id"`
	CorporateId  primitive.ObjectID `bson:"corporate_id"`
	Name         string             `bson:"name,omitempty"`
	ClientId     string             `bson:"client_id,omitempty"`
	ClientSecret string             `bson:"client_secret,omitempty"`
	Revoke       bool               `bson:"revoke"`
	ExpiredAt    time.Time          `bson:"expired_at,omitempty"`
	CreatedAt    time.Time          `bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at,omitempty"`
}
