package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Rating model
type Profile struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	DOB       time.Time          `json:"dob" bson:"dob,omitempty"`
	User      primitive.ObjectID `json:"user" bson:"user"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at" bson:"deleted_at,omitempty"`
}
