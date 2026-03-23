package services

import (
	"errors"
	"time"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/dtos"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/models"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/repositories"
	"github.com/google/uuid"
)

type ExpenseService struct {
	expenseRepository repositories.ExpenseRepository
}

func NewExpenseService(expenseRepository repositories.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		expenseRepository: expenseRepository,
	}
}

func (expenseService *ExpenseService) List(userID uuid.UUID) ([]dtos.ExpenseResponse, error) {
	return expenseService.expenseRepository.List(userID)
}

func (expenseService *ExpenseService) GetExpenseByIDAndUserID(id uuid.UUID, userId uuid.UUID) (*dtos.ExpenseResponse, error) {
	return expenseService.expenseRepository.GetExpenseByIdAndUserId(id, userId)
}

func (expenseService *ExpenseService) CreateExpense(newExpense dtos.ExpenseCreation) error {
	if newExpense.Category == "" {
		return errors.New("Category required")
	}

	if newExpense.Amount == 0 {
		return errors.New("Amount required")
	}

	var expense models.Expense

	expense.Amount = newExpense.Amount
	expense.UserId = newExpense.UserID
	expense.Category = newExpense.Category
	expense.Description = newExpense.Description
	expense.CreatedAt = time.Now()
	expense.UpdatedAt = time.Now()

	return expenseService.expenseRepository.CreateExpense(expense)
}
