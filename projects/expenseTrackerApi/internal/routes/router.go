package routes

import (
	"net/http"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/auth"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/handlers"
)

type Handlers struct {
	Auth    *auth.AuthHandler
	Expense *handlers.ExpenseHandler
}

func RegisterRouter(handlers *Handlers) http.Handler {
	mux := http.NewServeMux()

	RegisterAuthRoutes(mux, handlers.Auth)
	RegisterExpensesRoutes(mux, handlers.Expense)

	mux.Handle("/expenses", auth.JWTMiddleware(http.HandlerFunc(handlers.Expense.List)))

	return mux
}
