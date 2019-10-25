package models

import (
	"amovieplex-backend/src/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	PhoneNumber string             `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	Password    string             `json:"password" bson:"password"`
	Role        string             `json:"role" bson:"role"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt   time.Time          `json:"deleted_at" bson:"deleted_at,omitempty"`
}

func (u *User) FromJSON(json common.JSON) {
	u.ID, _ = primitive.ObjectIDFromHex(json["_id"].(string))
	u.Email = json["email"].(string)
	u.PhoneNumber = json["phone_number"].(string)
	u.Role = json["role"].(string)
}
