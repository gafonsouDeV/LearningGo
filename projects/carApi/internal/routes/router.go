package routes

import (
	"fmt"
	"net/http"

	"github.com/gafonsouDeV/LearningGo/projects/carApi/internal/middleware"
)

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	RegisterCarRoutes(mux)

	loggedRouter := middleware.Logging(mux)

	return loggedRouter
}
