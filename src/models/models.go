package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Genre model
type Genre struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at" bson:"deleted_at,omitempty"`
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
	ID          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string               `json:"title" bson:"title"`
	Plot        string               `json:"plot" bson:"plot"`
	Director    string               `json:"director" bson:"director"`
	Starring    string               `json:"starring" bson:"starring"`
	ReleaseDate time.Time            `json:"release_date" bson:"release_date"`
	RunningTime int                  `json:"running_time" bson:"running_time"`
	Genres      []primitive.ObjectID `json:"genres" bson:"genres"`
	Rating      primitive.ObjectID   `json:"rating" bson:"rating"`
	CreatedAt   time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at" bson:"updated_at"`
	DeletedAt   time.Time            `json:"deleted_at" bson:"deleted_at,omitempty"`
}
