package repositories

import (
	"context"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/dtos"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ExpensePostgresRepository struct {
	db *pgxpool.Pool
}

func NewExpensePostgresRepository(db *pgxpool.Pool) *ExpensePostgresRepository {
	return &ExpensePostgresRepository{db: db}
}

func (expenseRepository *ExpensePostgresRepository) List(userID uuid.UUID) ([]dtos.ExpenseResponse, error) {
	rows, err := expenseRepository.db.Query(context.Background(), `
		SELECT id, description, amount, category, created_at
		FROM expenses
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var expenses []dtos.ExpenseResponse

	for rows.Next() {
		var expense dtos.ExpenseResponse

		err := rows.Scan(
			&expense.ID,
			&expense.Description,
			&expense.Amount,
			&expense.Category,
			&expense.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (expenseRepository *ExpensePostgresRepository) GetExpenseByIdAndUserId(id uuid.UUID, userID uuid.UUID) (*dtos.ExpenseResponse, error) {
	query := `
		SELECT id, description, amount, category, created_at
		FROM expenses
		WHERE id = $1
	`

	var expense dtos.ExpenseResponse

	err := expenseRepository.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&expense.ID,
		&expense.Description,
		&expense.Category,
		&expense.Amount,
		&expense.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &expense, nil
}

func (expenseRepository *ExpensePostgresRepository) CreateExpense(expense models.Expense) error {
	query := `
		INSERT INTO expenses (user_id, description, category, amount, created_at, updated_at)
		VALUES ($1, $2, $3, $4 ,$5, $6)
	`

	_, err := expenseRepository.db.Exec(
		context.Background(),
		query,
		expense.UserId,
		expense.Description,
		expense.Category,
		expense.Amount,
		expense.CreatedAt,
		expense.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
