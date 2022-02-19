package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type image struct {
	Src   string `bson:"src"`
	Title string `bson:"title"`
}

type userInfo struct {
	Firstname string `bson:"firstname"`
	Lastname  string `bson:"lastname"`
	Phone     string `bson:"phone"`
	Image     image  `bson:"image"`
}

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username     string             `bson:"username,omitempty" json:"username,omitempty"`
	Password     string             `bson:"password,omitempty" json:"password,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	IsOnline     bool               `bson:"is_online" json:"is_online,omitempty"`
	Status       bool               `bson:"status" json:"status,omitempty"`
	LoginAttempt int                `bson:"login_attempt,omitempty" json:"login_attempt,omitempty"`
	UserInfo     userInfo           `bson:"user_info" json:"user_info,omitempty"`
	Roles        []Role             `bson:"role_id" json:"role_id,omitempty"`
	CreatedAt    time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
