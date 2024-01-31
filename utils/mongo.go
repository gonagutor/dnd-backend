package utils

import (
	"context"
	"log"
	"os"

	"github.com/TwiN/go-color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ClassCollection *mongo.Collection

func SetupMongoConnection() {
	clientOptions := options.Client().ApplyURI("mongodb://" + os.Getenv("MONGO_USERNAME") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(color.InRed("Could not connect to mongo: ") + err.Error())
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(color.InRed("Could not connect to mongo: ") + err.Error())
	}

	ClassCollection = client.Database("dnd").Collection("class")
}
