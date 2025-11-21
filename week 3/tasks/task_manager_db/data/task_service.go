package data

import (
	"context"
	"errors"
	"task_manager_db/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var taskCollection *mongo.Collection
var lastTaskID int

func InitTaskService(db *mongo.Database) {
	taskCollection = db.Collection("tasks")

	opts := options.FindOne().SetSort(bson.D{{"id", -1}})
	var lastTask models.Task
	err := taskCollection.FindOne(context.TODO(), bson.D{}, opts).Decode(&lastTask)
	if err == nil {
		lastTaskID = lastTask.ID
	} else {
		lastTaskID = 0
	}
}

// inserts a new task
func CreateTask(task models.Task) (models.Task, error) {
	lastTaskID++
	task.ID = lastTaskID
	task.Status = "pending"
	task.CreatedAt = time.Now()

	_, err := taskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

// UpdateTask by  ID
func UpdateTask(id int, update models.Task) (models.Task, error) {
	filter := bson.M{"id": id}
	updateBson := bson.M{
		"$set": bson.M{
			"title":       update.Title,
			"description": update.Description,
			"due_date":    update.DueDate,
			"status":      update.Status,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	res := taskCollection.FindOneAndUpdate(context.TODO(), filter, updateBson, opts)

	var updatedTask models.Task
	if err := res.Decode(&updatedTask); err != nil {
		return models.Task{}, errors.New("task not found")
	}

	return updatedTask, nil
}

// DeleteTask deletes a task by ID
func DeleteTask(id int) error {
	res, err := taskCollection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

// GetAllTasks retrieves all tasks
func GetAllTasks() ([]models.Task, error) {
	cursor, err := taskCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var tasks []models.Task
	if err := cursor.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTaskByID  by integer ID
func GetTaskByID(id int) (models.Task, error) {
	var task models.Task
	err := taskCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&task)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}
