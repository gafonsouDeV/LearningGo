package models

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Brand     string    `json:"brand"`
	Model     string    `json:"model"`
	Plate     string    `json:"plate"`
	Year      uint16    `json:"year"`
	CreatedAt time.Time `json:"created_at"`
}
