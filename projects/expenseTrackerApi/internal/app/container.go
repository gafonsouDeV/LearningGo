package app

import (
	"log"
	"log/slog"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/auth"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/db"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/handlers"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/logger"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/repositories"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	AuthHandler    *auth.AuthHandler
	ExpenseHandler *handlers.ExpenseHandler
	Pool           *pgxpool.Pool
	Logger         *slog.Logger
}

func NewContainer() *Container {
	appLogger := logger.New()
	pool, err := db.NewPostgresPool()
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	authRepo := repositories.NewUserPostgresRepository(pool, appLogger)
	authSvc := auth.NewAuthService(authRepo, appLogger)
	authHandler := auth.NewAuthHandler(authSvc)

	expenseRepo := repositories.NewExpensePostgresRepository(pool)
	expenseService := services.NewExpenseService(expenseRepo)
	expenseHandler := handlers.NewExpenseHandler(expenseService)

	return &Container{
		AuthHandler:    authHandler,
		ExpenseHandler: expenseHandler,
		Pool:           pool,
		Logger:         appLogger,
	}
}

func (c *Container) Close() {
	if c.Pool != nil {
		c.Pool.Close()
	}
}
