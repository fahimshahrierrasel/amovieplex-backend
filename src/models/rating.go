package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Rating model
type Rating struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	AgeLimit  int                `json:"age_limit" bson:"age_limit"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at" bson:"deleted_at,omitempty"`
}

