package data

import "time"

// Genre model
type Genre struct {
	Name string
}

// Rating model
type Rating struct {
	Name     string
	AgeLimit int
	CreateAt time.Time
}

// Movie model
type Movie struct {
	Title       string
	ReleaseYear string
	Genre       string
}

// Name model
type Name struct {
	FirstName string
}
