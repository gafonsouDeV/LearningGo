package routes

import (
	"net/http"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/auth"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/handlers"
)

func RegisterExpensesRoutes(mux *http.ServeMux, expenseHandler *handlers.ExpenseHandler) {
	expenseMux := http.NewServeMux()
	expenseMux.Handle("GET /", auth.JWTMiddleware(http.HandlerFunc(expenseHandler.List)))
	expenseMux.Handle("GET /{id}", auth.JWTMiddleware(http.HandlerFunc(expenseHandler.GetExpenseByIdAndUserId)))
	expenseMux.Handle("POST /create", auth.JWTMiddleware(http.HandlerFunc(expenseHandler.CreateExpense)))
	expenseMux.Handle("PUT /{id}", auth.JWTMiddleware(http.HandlerFunc(expenseHandler.UpdateExpense)))
	expenseMux.Handle("DELETE /{id}", auth.JWTMiddleware(http.HandlerFunc(expenseHandler.DeleteExpense)))

	mux.Handle("/expenses/", http.StripPrefix("/expenses", expenseMux))
}
