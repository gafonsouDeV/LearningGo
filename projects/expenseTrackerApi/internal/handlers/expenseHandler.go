package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/auth"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/dtos"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/services"
	"github.com/google/uuid"
)

type ExpenseHandler struct {
	expenseService *services.ExpenseService
}

func NewExpenseHandler(expenseService *services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: expenseService,
	}
}

func (expHandler *ExpenseHandler) List(resWriter http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resWriter, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userIDStr, ok := auth.FromContext(req.Context())

	if !ok {
		http.Error(resWriter, "unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(resWriter, "Invalid user id", http.StatusUnauthorized)
		return
	}

	expenses, err := expHandler.expenseService.List(userID)
	if err != nil {
		http.Error(resWriter, err.Error(), http.StatusInternalServerError)
	}

	resWriter.Header().Set("Content-Type", "application/json")
	response := struct {
		UserID   string                 `json:"user_id"`
		Message  string                 `json:"message"`
		Expenses []dtos.ExpenseResponse `json:"expenses"`
	}{
		UserID:   userIDStr,
		Message:  "expense list tracker",
		Expenses: expenses,
	}
	json.NewEncoder(resWriter).Encode(response)
}

func (expHandler *ExpenseHandler) GetExpenseByIdAndUserId(resWriter http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resWriter, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr, ok := auth.FromContext(req.Context())
	if !ok {
		http.Error(resWriter, "unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(resWriter, "Invalid user ID", http.StatusUnauthorized)
		return
	}

	expenseIDStr := req.PathValue("id")
	if expenseIDStr == "" {
		http.Error(resWriter, "Missing expense ID", http.StatusBadRequest)
		return
	}

	expenseId, err := uuid.Parse(expenseIDStr)
	if err != nil {
		http.Error(resWriter, "Invalid expense ID", http.StatusBadRequest)
		return
	}

	expense, err := expHandler.expenseService.GetExpenseByIDAndUserID(expenseId, userID)
	if err != nil {
		http.Error(resWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	if expense == nil {
		http.Error(resWriter, "expense not found", http.StatusNotFound)
		return
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(resWriter).Encode(map[string]*dtos.ExpenseResponse{
		"expense": expense,
	})
}

func (expHandler *ExpenseHandler) CreateExpense(resWriter http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(resWriter, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr, ok := auth.FromContext(req.Context())
	if !ok {
		http.Error(resWriter, "unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(resWriter, "Invalid user ID", http.StatusUnauthorized)
		return
	}

	var newExpense dtos.ExpenseCreation
	newExpense.UserID = userID
	err = json.NewDecoder(req.Body).Decode(&newExpense)
	if err != nil {
		http.Error(resWriter, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = expHandler.expenseService.CreateExpense(newExpense)
	if err != nil {
		http.Error(resWriter, err.Error(), http.StatusBadRequest)
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(resWriter).Encode("Expense created")
}

func (expHandler *ExpenseHandler) UpdateExpense(resWriter http.ResponseWriter, req http.Request) {
	if req.Method != http.MethodPut {
		http.Error(resWriter, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr, ok := auth.FromContext(req.Context())
	if !ok {
		http.Error(resWriter, "unauthorized", http.StatusUnauthorized)
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(resWriter, "Invalid user ID", http.StatusUnauthorized)
		return
	}

	expenseIDStr := req.PathValue("id")
	if expenseIDStr == "" {
		http.Error(resWriter, "Missing expense ID", http.StatusBadRequest)
		return
	}

	expenseID, err := uuid.Parse(expenseIDStr)
	if err != nil {
		http.Error(resWriter, "Missing expense id", http.StatusBadRequest)
	}

	var updatedExpense dtos.ExpenseUpdate
	updatedExpense.UserID = userID

	err = json.NewDecoder(req.Body).Decode(&updatedExpense)
	if err != nil {
		http.Error(resWriter, "Invalid body", http.StatusBadRequest)
		return
	}

	err = expHandler.expenseService.UpdateExpense(expenseID, updatedExpense)
	if err != nil {
		http.Error(resWriter, err.Error(), http.StatusBadRequest)
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(resWriter).Encode("Expense updated")
}
