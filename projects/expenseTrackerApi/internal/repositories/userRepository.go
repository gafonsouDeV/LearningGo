package repositories

import (
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/dtos"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/models"
)

type UserRepository interface {
	GetAllUsers() ([]dtos.UserResponse, error)
	CreateUser(models.User) error
	GetUserByEmail(string) (*dtos.UserResponse, error)
	GetAuthUserByEmail(string) (*models.User, error)
}
