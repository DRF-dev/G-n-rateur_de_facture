package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB connecte à notre base de donnée
func DB() (*mongo.Collection, context.CancelFunc, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	collection := client.Database("downloadfiles").Collection("files")
	if err != nil {
		return collection, cancel, err
	}

	return collection, cancel, nil
}

//ConnectToDB permet de se connecter à notre base de donnée
func ConnectToDB() (*mongo.Collection, context.CancelFunc) {
	collection, cancel, err := DB()
	if err != nil {
		log.Fatalf("Error connection to database: %v\n", err)
	}
	return collection, cancel
}
