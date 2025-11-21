package main

import (
	"context"
	"log"
	"task_manager_db/data"
	"task_manager_db/router"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoDB(uri, dbName string) *mongo.Database {
	clientOpts := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("Connected to MongoDB")
	return client.Database(dbName)
}

func main() {
	db := initMongoDB("mongodb://localhost:27017", "task_manager_db")

	// Initialize services with MongoDB collections
	data.InitUserService(db)
	data.InitTaskService(db)

	r := router.SetupRouter()
	r.Run(":8080")
}
