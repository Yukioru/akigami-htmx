package models

import (
	"context"
	"strconv"
	"time"

	"akigami.co/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateMethodsInput struct {
	Ctx        *fiber.Ctx
	Collection string
}

type schemas interface {
	UserSchema
}

type Model[T schemas] struct {
	Ctx        *fiber.Ctx
	Collection string
	Schema     T
}

func (m Model[T]) FindById(id primitive.ObjectID) T {
	timer := time.Now()

	result := T{}
	db.DB.Collection(m.Collection).FindOne(context.TODO(), bson.D{{
		Key: "_id", Value: id,
	}}).Decode(&result)

	m.Ctx.Response().Header.Add("Server-Timing", m.Collection+"_FindById;dur="+strconv.FormatInt(time.Since(timer).Milliseconds(), 10))

	return result
}

func (m Model[T]) FindOne(params ...bson.D) T {
	timer := time.Now()
	filter := bson.D{}
	if len(params) > 0 {
		filter = params[0]
	}

	result := T{}
	db.DB.Collection(m.Collection).FindOne(context.TODO(), filter).Decode(&result)

	m.Ctx.Response().Header.Add("Server-Timing", m.Collection+"_FindOne;dur="+strconv.FormatInt(time.Since(timer).Milliseconds(), 10))

	return result
}

func (m Model[T]) Find(params ...bson.D) []T {
	timer := time.Now()
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

	m.Ctx.Response().Header.Add("Server-Timing", m.Collection+"_Find;dur="+strconv.FormatInt(time.Since(timer).Milliseconds(), 10))

	return results
}

func CreateMethods[T schemas](params CreateMethodsInput) Model[T] {
	model := Model[T]{
		Ctx:        params.Ctx,
		Collection: params.Collection,
	}

	return model
}

func Get[T schemas](c *fiber.Ctx, collection string) Model[T] {
	return CreateMethods[T](CreateMethodsInput{
		Ctx:        c,
		Collection: collection,
	})
}
