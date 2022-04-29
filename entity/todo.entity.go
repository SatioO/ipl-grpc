package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoEntity struct {
	ID          primitive.ObjectID `bson:"id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Completed   bool               `bson:"completed"`
}
