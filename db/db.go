package db

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = context.TODO()
var Client *mongo.Client
var DB *mongo.Database

func Init() {
	mongoURI := os.Getenv("DB_URI")
	mongoDBName := os.Getenv("DB_NAME")
	if mongoURI == "" {
		log.Fatal("You must set your 'DB_URI' environment variable.")
	}
	if mongoDBName == "" {
		log.Fatal("You must set your 'DB_NAME' environment variable.")
	}

	var err error
	Client, err = mongo.Connect(Ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	DB = Client.Database("htmx")
}
