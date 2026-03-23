package routes

import (
	"net/http"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/auth"
)

func RegisterAuthRoutes(mux *http.ServeMux, authHandler *auth.AuthHandler) {
	authMux := http.NewServeMux()
	authMux.HandleFunc("/register", authHandler.Register)
	authMux.HandleFunc("/login", authHandler.Login)

	mux.Handle("/auth/", http.StripPrefix("/auth", authMux))
}
