package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/app"
	"github.com/gafonsouDeV/LearningGo/projects/expenseTrackerApi/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}
	container := app.NewContainer()
	defer container.Close()
	router := routes.RegisterRouter(&routes.Handlers{
		Auth:    container.AuthHandler,
		Expense: container.ExpenseHandler,
	})

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", router)
}
