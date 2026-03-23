package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gafonsouDeV/LearningGo/projects/carApi/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	router := routes.RegisterRoutes()

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", router)
}
