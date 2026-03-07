package services

import (
	"errors"
	"time"

	"github.com/gafonsouDeV/LearningGo/projects/carApi/internal/models"
	"github.com/gafonsouDeV/LearningGo/projects/carApi/internal/repositories"
	"github.com/google/uuid"
)

type CarService struct {
	repo repositories.CarRepository
}

func (carService *CarService) GetCarByPlate(plate string) (*models.Car, error) {
	return carService.repo.GetCarByPlate(plate)

}

func NewCarService(carRepository repositories.CarRepository) *CarService {
	return &CarService{
		repo: carRepository,
	}
}
func (carService *CarService) GetAllCars() ([]models.Car, error) {
	cars, err := carService.repo.GetAllCars()

	return cars, err
}

func (carService *CarService) CreateCar(car models.Car) error {
	if car.Brand == "" {
		return errors.New("brand required")
	}

	existing, _ := carService.repo.GetCarByPlate(car.Plate)

	if existing != nil {
		return errors.New("car already exists")
	}

	car.ID = uuid.New()
	car.CreatedAt = time.Now()

	return carService.repo.CreateCar(car)
}
