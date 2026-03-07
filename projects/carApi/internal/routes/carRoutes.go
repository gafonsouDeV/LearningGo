package routes

import (
	"log"
	"net/http"

	"github.com/gafonsouDeV/LearningGo/projects/carApi/internal/database"
	"github.com/gafonsouDeV/LearningGo/projects/carApi/internal/handlers"
	"github.com/gafonsouDeV/LearningGo/projects/carApi/internal/repositories"
	"github.com/gafonsouDeV/LearningGo/projects/carApi/internal/services"
)

func RegisterCarRoutes(mux *http.ServeMux) {
	pool, err := database.NewPostgresPool()
	if err != nil {
		log.Fatal(err)
	}

	repo := repositories.NewCarPostgresRepository(pool)
	service := services.NewCarService(repo)
	handler := handlers.NewCarHandler(*service)

	mux.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetCars(w, r)
		case http.MethodPost:
			handler.CreateCar(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/cars/", handler.GetCarByPlate)
}
