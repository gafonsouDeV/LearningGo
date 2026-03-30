package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/auth"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/dtos"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/errors"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/services"
	"github.com/google/uuid"
)

type ExpenseHandler struct {
	expenseService *services.ExpenseService
}

func NewExpenseHandler(expenseService *services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{expenseService: expenseService}
}

func getUserIDFromContext(r *http.Request) (uuid.UUID, error) {
	userIDStr, ok := auth.FromContext(r.Context())
	if !ok || userIDStr == "" {
		return uuid.Nil, errors.Unauthorized("missing_user_id", "missing user id", nil)
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, errors.BadRequest("invalid_user_id", "invalid user id format", err)
	}

	return userID, nil
}

func (expenseHandler *ExpenseHandler) List(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	expenses, err := expenseHandler.expenseService.List(userID)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	errors.WriteJSON(w, http.StatusOK, expenses)
}

func (expenseHandler *ExpenseHandler) GetExpenseByIdAndUserId(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	expenseID := r.PathValue("id")
	if expenseID == "" {
		errors.WriteError(w, errors.BadRequest("missing_expense_id", "missing expense id", nil))
		return
	}

	id, err := uuid.Parse(expenseID)
	if err != nil {
		errors.WriteError(w, errors.BadRequest("invalid_expense_id", "invalid expense id format", err))
		return
	}

	expense, err := expenseHandler.expenseService.GetExpenseByIDAndUserID(id, userID)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	errors.WriteJSON(w, http.StatusOK, expense)
}

func (expenseHandler *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	_, err := getUserIDFromContext(r)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	var req dtos.ExpenseCreation
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, errors.BadRequest("invalid_json", "invalid request body", err))
		return
	}

	err = expenseHandler.expenseService.CreateExpense(req)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	errors.WriteJSON(w, http.StatusCreated, map[string]string{"message": "expense created successfully"})
}

func (expenseHandler *ExpenseHandler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	expenseID := r.PathValue("id")
	if expenseID == "" {
		errors.WriteError(w, errors.BadRequest("missing_expense_id", "missing expense id", nil))
		return
	}

	id, err := uuid.Parse(expenseID)
	if err != nil {
		errors.WriteError(w, errors.BadRequest("invalid_expense_id", "invalid expense id format", err))
		return
	}

	var req dtos.ExpenseUpdate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, errors.BadRequest("invalid_json", "invalid request body", err))
		return
	}

	req.UserID = userID

	err = expenseHandler.expenseService.UpdateExpense(id, req)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	errors.WriteJSON(w, http.StatusOK, map[string]string{"message": "expense updated successfully"})
}

func (expenseHandler *ExpenseHandler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	expenseID := r.PathValue("id")
	if expenseID == "" {
		errors.WriteError(w, errors.BadRequest("missing_expense_id", "missing expense id", nil))
		return
	}

	id, err := uuid.Parse(expenseID)
	if err != nil {
		errors.WriteError(w, errors.BadRequest("invalid_expense_id", "invalid expense id format", err))
		return
	}

	err = expenseHandler.expenseService.DeleteExpense(id, userID)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	errors.WriteJSON(w, http.StatusOK, map[string]string{"message": "expense deleted successfully"})
}
