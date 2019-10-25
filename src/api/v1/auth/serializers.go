package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserSerializer is a single genre serializer
type UserSerializer struct {
	User
}

// SingleUserResponse is the structure of the return response
type SingleUserResponse struct {
	ID          primitive.ObjectID `json:"_id"`
	Email       string             `json:"email"`
	PhoneNumber string             `json:"phone_number"`
	Role        string             `json:"role"`
}

func (us *UserSerializer) Response() SingleUserResponse {
	return SingleUserResponse{
		ID:          us.ID,
		Email:       us.Email,
		PhoneNumber: us.PhoneNumber,
		Role:        us.Role,
	}
}
