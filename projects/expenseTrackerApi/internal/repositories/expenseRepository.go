package repositories

import (
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/dtos"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/models"
	"github.com/google/uuid"
)

type ExpenseRepository interface {
	List(uuid.UUID) ([]dtos.ExpenseResponse, error)
	GetExpenseByIdAndUserId(uuid.UUID, uuid.UUID) (*dtos.ExpenseResponse, error)
	CreateExpense(models.Expense) error
	UpdateExpense(uuid.UUID, dtos.ExpenseUpdate) error
	DeleteExpense(uuid.UUID, uuid.UUID) error
}
