package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/dtos"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/errors"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (authHandler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dtos.UserCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, errors.BadRequest("invalid_json", "invalid request body", err))
		return
	}

	err := authHandler.authService.Register(req)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	errors.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user registered successfully"})
}

func (authHandler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dtos.UserCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, errors.BadRequest("invalid_json", "invalid request body", err))
		return
	}

	token, err := authHandler.authService.Login(req)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	errors.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}
