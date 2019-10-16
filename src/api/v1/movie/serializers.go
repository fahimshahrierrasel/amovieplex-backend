package movie

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// SimpleSerializer is the movie serializer
type SimpleSerializer struct {
	Movie
}

// SimpleResponse structure the return response
type SimpleResponse struct {
	ID          primitive.ObjectID   `json:"_id"`
	Title       string               `json:"title"`
	Plot        string               `json:"plot"`
	Director    string               `json:"director"`
	Starring    string               `json:"starring"`
	ReleaseDate time.Time            `json:"release_date"`
	RunningTime int                  `json:"running_time"`
	Genres      []primitive.ObjectID `json:"genres"`
	Rating      primitive.ObjectID   `json:"rating"`
}

func (ss *SimpleSerializer) Response() SimpleResponse {
	return SimpleResponse{
		ID:          ss.ID,
		Title:       ss.Title,
		Plot:        ss.Plot,
		Director:    ss.Director,
		Starring:    ss.Starring,
		ReleaseDate: ss.ReleaseDate,
		RunningTime: ss.RunningTime,
		Genres:      ss.Genres,
		Rating:      ss.Rating,
	}
}

// Serializer is the all movies serializer
type Serializer struct {
	Movies []Movie
}

func (s *Serializer) Response() []SimpleResponse {
	var response []SimpleResponse

	for _, movie := range s.Movies {
		serializer := SimpleSerializer{movie}
		response = append(response, serializer.Response())
	}
	return response
}
