package repositories

import "github.com/gafonsouDeV/LearningGo/projects/carApi/internal/models"

type CarMemoryRepository struct {
	cars []models.Car
}

func NewCarMemoryRepository() *CarMemoryRepository {
	return &CarMemoryRepository{
		cars: []models.Car{},
	}
}

func (carRepository *CarMemoryRepository) CreateCar(car models.Car) error {
	carRepository.cars = append(carRepository.cars, car)
	return nil
}

func (carRepository *CarMemoryRepository) GetAllCars() ([]models.Car, error) {
	return carRepository.cars, nil
}

func (carRepository *CarMemoryRepository) GetCarByPlate(plate string) (*models.Car, error) {
	for i := range carRepository.cars {
		if carRepository.cars[i].Plate == plate {
			return &carRepository.cars[i], nil
		}
	}

	return nil, nil
}
