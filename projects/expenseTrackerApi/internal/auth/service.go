package auth

import (
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/dtos"
	apperrors "github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/errors"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/models"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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
		return apperrors.BadRequest("email_required", "email required", nil)
	}

	if user.Password == "" {
		return apperrors.BadRequest("password_required", "password required", nil)
	}

	existing, err := authService.userRepo.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		authService.logger.Error("register_check_existing_user_failed", "email", user.Email, "error", err)
		return apperrors.Internal(err)
	}

	if existing != nil {
		authService.logger.Warn("register_failed_email_exists", "email", user.Email)
		return apperrors.Conflict("email_already_registered", "that email is already registered", nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		authService.logger.Error("register_hash_password_failed", "email", user.Email, "error", err)
		return apperrors.Internal(err)
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

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return apperrors.Conflict("email_already_registered", "that email is already registered", err)
		}

		return apperrors.Internal(err)
	}

	authService.logger.Info("user_registered", "user_id", newUser.ID, "email", newUser.Email, "role", newUser.Role)
	return nil
}

func (authService *AuthService) Login(user dtos.UserCreateRequest) (string, error) {
	if user.Email == "" {
		return "", apperrors.BadRequest("email_required", "email required", nil)
	}

	if user.Password == "" {
		return "", apperrors.BadRequest("password_required", "password required", nil)
	}

	authUser, err := authService.userRepo.GetAuthUserByEmail(user.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			authService.logger.Warn("login_failed_invalid_email", "email", user.Email)
			return "", apperrors.Unauthorized("invalid_credentials", "invalid credentials", nil)
		}

		authService.logger.Error("login_get_user_failed", "email", user.Email, "error", err)
		return "", apperrors.Internal(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(authUser.PasswordHash), []byte(user.Password))
	if err != nil {
		authService.logger.Warn("login_failed_invalid_email", "email", user.Email)
		return "", apperrors.Unauthorized("invalid_credentials", "invalid credentials", nil)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": authUser.ID,
		"email":   authUser.Email,
		"role":    authUser.Role,
		"expires": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(authService.jwtSecret))
	if err != nil {
		authService.logger.Error("login_sign_token_failed", "user_id", authUser.ID, "error", err)
		return "", apperrors.Internal(err)
	}
	authService.logger.Info("loggin_success", "user_id", authUser.ID, "email", authUser.Email, "role", authUser.Role)
	return tokenString, nil
}
