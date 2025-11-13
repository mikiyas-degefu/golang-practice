package data

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"task_manager_db/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client       *mongo.Client
	db           *mongo.Database
	tasksColl    *mongo.Collection
	countersColl *mongo.Collection
)

const (
	defaultMongoURI = "mongodb://localhost:27017"
	defaultDBName   = "taskdb"
	taskIDCounter   = "taskid"
)

// InitDB initializes MongoDB client and collections. It reads env vars:
// MONGODB_URI (default mongodb://localhost:27017) and MONGODB_DB (default taskdb).
func InitDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = defaultMongoURI
	}
	dbName := os.Getenv("MONGODB_DB")
	if dbName == "" {
		dbName = defaultDBName
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	clientOpts := options.Client().ApplyURI(uri)
	client, err = mongo.Connect(ctx, clientOpts)
	if err != nil {
		return fmt.Errorf("mongo connect error: %w", err)
	}

	// Ping
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("mongo ping error: %w", err)
	}

	db = client.Database(dbName)
	tasksColl = db.Collection("tasks")
	countersColl = db.Collection("counters")

	// Ensure counter doc exists (upsert will create on first use; we don't force create here)
	return nil
}

// GetAll returns all tasks
func GetAll() []models.Task {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := tasksColl.Find(ctx, bson.M{})
	if err != nil {
		return []models.Task{}
	}
	defer cursor.Close(ctx)

	var list []models.Task
	if err := cursor.All(ctx, &list); err != nil {
		return []models.Task{}
	}
	return list
}

// GetByID returns a task by id
func GetByID(id int) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var t models.Task
	err := tasksColl.FindOne(ctx, bson.M{"id": id}).Decode(&t)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("not found")
		}
		return models.Task{}, err
	}
	return t, nil
}

// getNextID atomically increments and returns the next integer ID for tasks.
func getNextID(ctx context.Context) (int, error) {
	filter := bson.M{"_id": taskIDCounter}
	update := bson.M{"$inc": bson.M{"seq": 1}}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var res struct {
		Seq int `bson:"seq"`
	}

	err := countersColl.FindOneAndUpdate(ctx, filter, update, opts).Decode(&res)
	if err != nil {
		// If decode fails because the counter doesn't exist yet, create it manually
		if err == mongo.ErrNoDocuments {
			initDoc := bson.M{"_id": taskIDCounter, "seq": 1}
			if _, err := countersColl.InsertOne(ctx, initDoc); err != nil {
				return 0, err
			}
			return 1, nil
		}
		return 0, err
	}

	return res.Seq, nil
}

// getNextID atomically increments and returns the next integer ID for tasks.
// Create a new task
func Create(t models.Task) models.Task {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// generate sequential int id
	id, err := getNextID(ctx)
	if err != nil {
		// fallback: set to timestamp low bits
		id = int(time.Now().Unix())
	}
	t.ID = id

	res, err := tasksColl.InsertOne(ctx, t)
	if err == nil {
		// set the MongoID if returned
		if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
			t.MongoID = oid
		}
	}
	return t
}

// Update an existing task
func Update(id int, updated models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ensure existing
	var existing models.Task
	err := tasksColl.FindOne(ctx, bson.M{"id": id}).Decode(&existing)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("not found")
		}
		return models.Task{}, err
	}

	updated.ID = id
	filter := bson.M{"id": id}
	repl := bson.M{"$set": bson.M{
		"title":       updated.Title,
		"description": updated.Description,
		"due_date":    updated.DueDate,
		"status":      updated.Status,
		"id":          updated.ID,
	}}

	_, err = tasksColl.UpdateOne(ctx, filter, repl)
	if err != nil {
		return models.Task{}, err
	}
	// return the updated version
	err = tasksColl.FindOne(ctx, filter).Decode(&updated)
	if err != nil {
		return models.Task{}, err
	}
	return updated, nil
}

// Delete removes a task
func Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := tasksColl.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("not found")
	}
	return nil
}
