package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShortLink struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      primitive.ObjectID `bson:"user_id"`
	RefCode     string             `bson:"ref_code"`
	OriginalUrl string             `bson:"original_url"`
	ShortUrl    string             `bson:"short_url"`
	ExpiredAt   time.Time          `bson:"expired_at"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
