package models

import "time"

// Genre model
type Genre struct {
	Name string
}

// Rating model
type Rating struct {
	Name      string
	AgeLimit  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `bson:",omitempty"`
}

// Movie model
type Movie struct {
	Title       string
	ReleaseYear string
	Genre       string
}
