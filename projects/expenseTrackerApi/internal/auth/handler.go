package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/dtos"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (handler *AuthHandler) Register(resWriter http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(resWriter, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var userRequest dtos.UserCreateRequest
	err := json.NewDecoder(req.Body).Decode(&userRequest)
	if err != nil {
		http.Error(resWriter, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = handler.authService.Register(userRequest)

	if err != nil {
		http.Error(resWriter, err.Error(), http.StatusBadRequest)
		return
	}
	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(resWriter).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

func (handler *AuthHandler) Login(resWriter http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(resWriter, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginRequest dtos.UserCreateRequest
	err := json.NewDecoder(req.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(resWriter, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	token, err := handler.authService.Login(loginRequest)
	if err != nil {
		http.Error(resWriter, err.Error(), http.StatusBadRequest)
		return
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(resWriter).Encode(map[string]string{
		"token": token,
	})
}
