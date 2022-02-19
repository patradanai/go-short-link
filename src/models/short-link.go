package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShortLink struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserId      primitive.ObjectID `bson:"user_id" json:"user_id"`
	RefCode     string             `bson:"ref_code,omitempty" json:"ref_code,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	SlugTag     string             `bson:"slug_tag,omitempty" json:"slug_tag,omitempty"`
	OriginalUrl string             `bson:"original_url,omitempty" json:"original_url,omitempty"`
	ShortUrl    string             `bson:"short_url,omitempty" json:"short_url,omitempty"`
	ExpiredAt   time.Time          `bson:"expired_at,omitempty" json:"expired_at,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
