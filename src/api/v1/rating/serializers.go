package rating

import "go.mongodb.org/mongo-driver/bson/primitive"

// SimpleSerializer is the rating serializer
type SimpleSerializer struct {
	Rating
}

// SimpleResponse structure the return response
type SimpleResponse struct {
	ID       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	AgeLimit int                `json:"age_limit"`
}

func (ss *SimpleSerializer) Response() SimpleResponse {
	return SimpleResponse{
		ID:       ss.ID,
		Name:     ss.Name,
		AgeLimit: ss.AgeLimit,
	}
}

// Serializer is the all rating serializer
type Serializer struct {
	Ratings []Rating
}

func (s *Serializer) Response() []SimpleResponse {
	var response []SimpleResponse

	for _, rating := range s.Ratings {
		serializer := SimpleSerializer{rating}
		response = append(response, serializer.Response())
	}
	return response
}
