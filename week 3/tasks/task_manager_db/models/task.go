package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Task represents a task in the system. Keep `id` as an int to remain backward compatible
// with the previous JSON format. Mongo's native _id is stored in MongoID but omitted
// from JSON output.
type Task struct {
	MongoID     primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ID          int                `json:"id" bson:"id"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	DueDate     string             `json:"due_date,omitempty" bson:"due_date,omitempty"`
	Status      string             `json:"status" bson:"status" binding:"required"`
}
