package dtos

import (
	"time"

	"github.com/google/uuid"
)

type ExpenseResponse struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"string"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
}

type ExpenseCreation struct {
	UserID      uuid.UUID `json:"id"`
	Description string    `json:"string"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
}
