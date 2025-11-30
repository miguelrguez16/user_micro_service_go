package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB establishes a connection to the MongoDB database.
func ConnectDB() (db *mongo.Database, err error, isProduction bool) {
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	isProduction = os.Getenv("ENV") == "production"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err, isProduction
	}

	log.Println("Connected to MongoDB")
	return client.Database(dbName), nil, isProduction
}
