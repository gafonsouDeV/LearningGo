package auth

import (
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/dtos"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/models"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  repositories.UserRepository
	jwtSecret string
	logger    *slog.Logger
}

func NewAuthService(userRepository repositories.UserRepository, logger *slog.Logger) *AuthService {
	return &AuthService{
		userRepo:  userRepository,
		jwtSecret: os.Getenv("JWT_SECRET"),
		logger:    logger,
	}
}

func (authService *AuthService) Register(user dtos.UserCreateRequest) error {
	if user.Email == "" {
		return errors.New("email required")
	}

	if user.Password == "" {
		return errors.New("password required")
	}

	existing, _ := authService.userRepo.GetUserByEmail(user.Email)

	if existing != nil {
		authService.logger.Warn("register_failed_email_exists", "email", user.Email)
		return errors.New("that email is already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return errors.New("error hashing the password")
	}

	var newUser models.User

	newUser.ID = uuid.New()
	newUser.Email = user.Email
	newUser.PasswordHash = string(hashedPassword)
	newUser.Role = "user"
	newUser.CreatedAt = time.Now()

	err = authService.userRepo.CreateUser(newUser)
	if err != nil {
		authService.logger.Error("register_user_failed", "email", user.Email, "error", err)
		return err
	}

	authService.logger.Info("user_registered", "user_id", newUser.ID, "email", newUser.Email, "role", newUser.Role)
	return nil
}

func (authService *AuthService) Login(user dtos.UserCreateRequest) (string, error) {
	authUser, err := authService.userRepo.GetAuthUserByEmail(user.Email)
	if err != nil {
		authService.logger.Warn("login_failed_invalid_email", "email", user.Email)
		return "", errors.New("invalid email credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(authUser.PasswordHash), []byte(user.Password))
	if err != nil {
		authService.logger.Warn("login_failed_invalid_password", "email", user.Email)
		return "", errors.New("invalid password credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": authUser.ID,
		"email":   authUser.Email,
		"role":    authUser.Role,
		"expires": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(authService.jwtSecret))
	if err != nil {
		return "", err
	}
	authService.logger.Info("loggin_success", "user_id", authUser.ID, "email", authUser.Email, "role", authUser.Role)
	return tokenString, nil
}
