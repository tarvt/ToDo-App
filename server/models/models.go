package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDoList struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id ,omitempty"`
	Status bool               `json:"status,omitempty"`
	Task   string             `json:"task,omitempty"`
}
