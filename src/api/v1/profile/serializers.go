package profile

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProfileSerializer struct {
	Profile
}

type SingleProfileResponse struct {
	ID        primitive.ObjectID `json:"_id"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	DOB       time.Time          `json:"dob"`
}

func (ps *ProfileSerializer) Response() SingleProfileResponse {
	return SingleProfileResponse{
		ID:        ps.ID,
		FirstName: ps.FirstName,
		LastName:  ps.LastName,
		DOB:       ps.DOB,
	}
}
