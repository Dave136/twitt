package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection = connection()
var clientOptions *options.ClientOptions

func connection() *mongo.Client {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	mongoUri := os.Getenv("MONGO_URI")

	log.Println(mongoUri)

	clientOptions = options.Client().ApplyURI(mongoUri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Successfully connected to database 🔥")

	return client
}

func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
