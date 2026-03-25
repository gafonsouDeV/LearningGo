package services

import (
	"errors"
	"log/slog"
	"time"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/dtos"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/models"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/repositories"
	"github.com/google/uuid"
)

type ExpenseService struct {
	expenseRepository repositories.ExpenseRepository
	logger            *slog.Logger
}

func NewExpenseService(expenseRepository repositories.ExpenseRepository, logger *slog.Logger) *ExpenseService {
	return &ExpenseService{
		expenseRepository: expenseRepository,
		logger:            logger,
	}
}

func (expenseService *ExpenseService) List(userID uuid.UUID) ([]dtos.ExpenseResponse, error) {
	expenses, err := expenseService.expenseRepository.List(userID)
	if err != nil {
		expenseService.logger.Error("get_expense_list_failed", "error", err)
		return nil, err
	}

	expenseService.logger.Info("get_expense_list_successfully", "user_id", userID)
	return expenses, nil
}

func (expenseService *ExpenseService) GetExpenseByIDAndUserID(id uuid.UUID, userId uuid.UUID) (*dtos.ExpenseResponse, error) {
	expense, err := expenseService.expenseRepository.GetExpenseByIdAndUserId(id, userId)

	if err != nil {
		expenseService.logger.Error("get_expense_failed", "error", err)
		return nil, err
	}

	expenseService.logger.Info("get_expense_successfully", "user_id", userId, "expense_id", id)
	return expense, nil
}

func (expenseService *ExpenseService) CreateExpense(newExpense dtos.ExpenseCreation) error {
	if newExpense.Category == "" {
		expenseService.logger.Warn("create_expense_failed_no_category", "user_id", newExpense.UserID)
		return errors.New("Category required")
	}

	if newExpense.Amount == 0 {
		expenseService.logger.Warn("create_expense_failed_no_amount", "user_id", newExpense.UserID)
		return errors.New("Amount required")
	}

	var expense models.Expense

	expense.Amount = newExpense.Amount
	expense.UserId = newExpense.UserID
	expense.Category = newExpense.Category
	expense.Description = newExpense.Description
	expense.CreatedAt = time.Now()
	expense.UpdatedAt = time.Now()

	err := expenseService.expenseRepository.CreateExpense(expense)

	if err != nil {
		expenseService.logger.Error("create_expense_failed", "user_id", expense.UserId, "error", err)
		return err
	}

	expenseService.logger.Info("create_expense_successfully", "user_id", expense.UserId, "expense_id", expense.ID)
	return nil
}

func (expenseService *ExpenseService) UpdateExpense(expenseId uuid.UUID, updatedExpense dtos.ExpenseUpdate) error {
	if updatedExpense.Category == "" {
		expenseService.logger.Warn("update_expense_failed_no_category", "user_id", updatedExpense.UserID, "expense_id", expenseId)
		return errors.New("Category required")
	}

	if updatedExpense.Amount == 0 {
		expenseService.logger.Warn("update_expense_failed_no_amount", "user_id", updatedExpense.UserID, "expense_id", expenseId)
		return errors.New("Amount required")
	}

	err := expenseService.expenseRepository.UpdateExpense(expenseId, updatedExpense)
	if err != nil {
		expenseService.logger.Error("update_expense_failed", "user_id", updatedExpense.UserID, "expense_id", expenseId)
		return err
	}

	expenseService.logger.Info("update_expense_successfully", "user_id", updatedExpense.UserID, "expense_id", expenseId)
	return nil
}
