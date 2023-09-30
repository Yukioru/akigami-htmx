package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSchema struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string
	Username string
	Avatar   string
}
