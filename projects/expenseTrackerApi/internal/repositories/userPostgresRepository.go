package repositories

import (
	"context"
	"log/slog"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/dtos"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPostgresRepository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewUserPostgresRepository(db *pgxpool.Pool, logger *slog.Logger) *UserPostgresRepository {
	return &UserPostgresRepository{
		db:     db,
		logger: logger,
	}
}

func (userRepository *UserPostgresRepository) GetAllUsers() ([]dtos.UserResponse, error) {
	rows, err := userRepository.db.Query(context.Background(), `
	SELECT id, email, created_at FROM users
	`)

	if err != nil {
		userRepository.logger.Error("db_get_all_users_failed", "error", err)
		return nil, err
	}

	defer rows.Close()

	var users []dtos.UserResponse

	for rows.Next() {
		var user dtos.UserResponse

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.CreatedAt,
		)

		if err != nil {
			userRepository.logger.Error("db_scan_user_failed", "error", err)
			return nil, err
		}

		users = append(users, user)
	}

	userRepository.logger.Debug("db_get_all_user_success", "count", len(users))
	return users, nil
}

func (userRepository *UserPostgresRepository) GetUserByEmail(email string) (*dtos.UserResponse, error) {
	query := `
		SELECT id, email, created_at 
		FROM users
		WHERE email = $1
	`
	var user dtos.UserResponse

	err := userRepository.db.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		userRepository.logger.Debug("db_get_user_by_email_not_found", "email", email)
		return nil, err
	}

	userRepository.logger.Debug("db_get_user_by_email_success", "email", email, "user_id", user.ID)
	return &user, nil
}

func (userRepository *UserPostgresRepository) GetAuthUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT * 
		FROM users
		WHERE email = $1
	`

	var user models.User
	err := userRepository.db.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		userRepository.logger.Debug("db_get_auth_user_by_email_not_found", "email", email)
		return nil, err
	}

	userRepository.logger.Debug("db_get_auth_user_by_email_success", "email", email, "user_id", user.ID)
	return &user, nil
}

func (userRepository *UserPostgresRepository) CreateUser(newUser models.User) error {
	query := `
		INSERT INTO users (id, email, password_hash, role, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := userRepository.db.Exec(
		context.Background(),
		query,
		newUser.ID,
		newUser.Email,
		newUser.PasswordHash,
		newUser.Role,
		newUser.CreatedAt,
	)
	if err != nil {
		userRepository.logger.Error("db_create_user_failed", "email", newUser.Email, "error", err)
		return err
	}

	userRepository.logger.Info("db_create_user_success", "user_id", newUser.ID, "email", newUser.Email, "role", newUser.Role)
	return nil
}
