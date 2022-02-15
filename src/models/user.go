package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type image struct {
	Src string `bson:"src"`
	Title string `bson:"title"`
}

type userInfo struct {
	Firstname string `bson:"firstname"`
	Lastname string `bson:"lastname"`
	Phone string `bson:"phone"`
	Image image `bson:"image"`
}


type User struct {
	Id primitive.ObjectID `bson:"_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	IsOnline bool `bson:"is_online"`
	Status bool `bson:"status"`
	LoginAttempt int `bson:"login_attempt"`
	UserInfo userInfo `bson:"user_info"`
	Roles []string `bson:"roles"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}