package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system.
type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name" binding:"required"`
	Email string             `bson:"email" json:"email" binding:"required,email"`
}
