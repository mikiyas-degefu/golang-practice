package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	MongoID     primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ID          int                `json:"id" bson:"id"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	DueDate     string             `json:"due_date,omitempty" bson:"due_date,omitempty"`
	Status      string             `json:"status" bson:"status" binding:"required"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
