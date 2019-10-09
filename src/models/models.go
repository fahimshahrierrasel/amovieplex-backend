package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Genre model
type Genre struct {
	Name string
}

// Rating model
type Rating struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	AgeLimit  int                `json:"age_limit" bson:"age_limit"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at" bson:"deleted_at,omitempty"`
}

// Movie model
type Movie struct {
	Title       string
	ReleaseYear string
	Genre       string
}
