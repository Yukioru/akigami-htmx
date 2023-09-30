package models

import (
	"context"

	"akigami.co/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type schemas interface {
	UserSchema
}

type Model[T schemas] struct {
	Collection string
	Schema     T
}

func (m Model[T]) FindById(id primitive.ObjectID) T {
	result := T{}
	db.DB.Collection(m.Collection).FindOne(context.TODO(), bson.D{{
		Key: "_id", Value: id,
	}}).Decode(&result)

	return result
}

func (m Model[T]) FindOne(params ...bson.D) T {
	filter := bson.D{}
	if len(params) > 0 {
		filter = params[0]
	}

	result := T{}
	db.DB.Collection(m.Collection).FindOne(context.TODO(), filter).Decode(&result)

	return result
}

func (m Model[T]) Find(params ...bson.D) []T {
	filter := bson.D{}
	if len(params) > 0 {
		filter = params[0]
	}

	cursor, err := db.DB.Collection(m.Collection).Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	results := []T{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func CreateMethods[T schemas](collection string) Model[T] {
	model := Model[T]{
		Collection: collection,
	}

	return model
}
