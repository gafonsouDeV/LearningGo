package models

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID          uuid.UUID `db:"id"`
	UserId      uuid.UUID `db:"user_id"`
	Description string    `db:"description"`
	Amount      float64   `db:"amount"`
	Category    string    `db:"category"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
