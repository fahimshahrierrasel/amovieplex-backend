package genre

import "go.mongodb.org/mongo-driver/bson/primitive"

// SimpleSerializer is a single genre serializer
type SimpleSerializer struct {
	Genre
}

// SimpleResponse is the structure of the return response
type SimpleResponse struct {
	ID   primitive.ObjectID `json:"_id"`
	Name string             `json:"name"`
}

func (ss *SimpleSerializer) Response() SimpleResponse {
	return SimpleResponse{
		ID:   ss.ID,
		Name: ss.Name,
	}
}

// Serializer is all genre serializer
type Serializer struct {
	Genres []Genre
}

func (s *Serializer) Response() []SimpleResponse {
	var response []SimpleResponse

	for _, genre := range s.Genres {
		serializer := SimpleSerializer{genre}
		response = append(response, serializer.Response())
	}
	return response
}
