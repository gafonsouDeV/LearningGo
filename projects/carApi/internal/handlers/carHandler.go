package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gafonsouDeV/LearningGo/projects/carApi/internal/models"
	"github.com/gafonsouDeV/LearningGo/projects/carApi/internal/services"
)

type CarHandler struct {
	service services.CarService
}

func NewCarHandler(carService services.CarService) *CarHandler {
	return &CarHandler{
		service: carService,
	}
}

func (handler *CarHandler) GetCars(resWriter http.ResponseWriter, req *http.Request) {
	cars, err := handler.service.GetAllCars()

	if err != nil {
		http.Error(resWriter, "Failed to retrieve cars", http.StatusInternalServerError)
	}

	json.NewEncoder(resWriter).Encode(cars)
}

func (handler *CarHandler) GetCarByPlate(resWriter http.ResponseWriter, req *http.Request) {
	plate := strings.TrimPrefix(req.URL.Path, "/cars/")

	car, err := handler.service.GetCarByPlate(plate)

	if err != nil {
		http.Error(resWriter, "Error retrieving the car", http.StatusInternalServerError)
		return
	}

	if car == nil {
		http.Error(resWriter, "Car not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(resWriter).Encode(car)
}

func (handler *CarHandler) CreateCar(resWriter http.ResponseWriter, req *http.Request) {
	var car models.Car

	err := json.NewDecoder(req.Body).Decode(&car)

	if err != nil {
		http.Error(resWriter, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	err = handler.service.CreateCar(car)

	if err != nil {
		http.Error(resWriter, err.Error(), http.StatusBadRequest)
		return
	}

	resWriter.WriteHeader(http.StatusCreated)

	json.NewEncoder(resWriter).Encode(car)
}

func (handler *CarHandler) UpdateCar(resWriter http.ResponseWriter, req *http.Request) {
	var updatedCar models.Car

	err := json.NewDecoder(req.Body).Decode(&updatedCar)
	if err != nil {
		http.Error(resWriter, "Invalid JSON body", http.StatusBadRequest)
		return
	}

}
