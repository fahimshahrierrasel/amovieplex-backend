package models

import "time"

// Genre model
type Genre struct {
	Name string
}

// Rating model
type Rating struct {
	Name      string    `json:"name" bson:"name"`
	AgeLimit  int       `json:"age_limit" bson:"age_limit"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at,omitempty"`
}

// Movie model
type Movie struct {
	Title       string
	ReleaseYear string
	Genre       string
}
