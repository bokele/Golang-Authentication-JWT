package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	first_name    *string            `json:"firsst_name" validate:"required, min=2, max=100"`
	last_name     *string            `json:"last_name" validate:"required, min=2, max=100"`
	password      *string            `json:"password" validate:"required, min=6"`
	email         *string            `json:"email" validate:"required, email"`
	phone         *string            `json:"phone" validate:"required, min=8, max=10"`
	Token         *string            `json:"token" validate:"required, min=2, max=100"`
	User_type     *string            `json:"user_type: validate:"required, eq=admin|eq=user"`
	refresh_token *string            `json:"refresh_token"`
	created_at    time.Time          `json:"created_at"`
	updated_at    time.Time          `json:"updated_at"`
	User_id       *string            `json:"User_id"`
}
