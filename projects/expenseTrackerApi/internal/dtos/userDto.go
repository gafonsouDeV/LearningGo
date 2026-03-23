package dtos

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
