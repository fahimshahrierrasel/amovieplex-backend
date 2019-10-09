package rating

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SingleRatingSerializer is the rating serializer
type SingleRatingSerializer struct {
	Rating
}

// SimpleRatingResponse structure the return response
type SimpleRatingResponse struct {
	ID       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	AgeLimit int                `json:"age_limit"`
}

// Response return the SimpleRatingResponse
func (self *SingleRatingSerializer) Response() SimpleRatingResponse {
	return SimpleRatingResponse{
		ID:       self.ID,
		Name:     self.Name,
		AgeLimit: self.AgeLimit,
	}
}

// RatingSerializer is the all rating serializer
type RatingSerializer struct {
	Ratings []Rating
}

func (self *RatingSerializer) Response() []SimpleRatingResponse {
	response := []SimpleRatingResponse{}

	for _, rating := range self.Ratings {
		serializer := SingleRatingSerializer{rating}
		response = append(response, serializer.Response())
	}
	return response
}
