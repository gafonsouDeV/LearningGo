package repositories

import "github.com/gafonsouDeV/LearningGo/projects/carApi/internal/models"

type CarRepository interface {
	GetAllCars() ([]models.Car, error)
	CreateCar(car models.Car) error
	GetCarByPlate(plate string) (*models.Car, error) // pointer is used to avoid copy structs and return nil
}
