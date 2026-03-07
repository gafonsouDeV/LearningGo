package main

import (
	"fmt"
	"net/http"

	"github.com/gafonsouDeV/LearningGo/projects/carApi/internal/routes"
)

func main() {
	router := routes.RegisterRoutes()

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", router)
}
