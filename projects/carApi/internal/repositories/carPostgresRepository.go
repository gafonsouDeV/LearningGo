package repositories

import (
	"context"

	"github.com/gafonsouDeV/LearningGo/projects/carApi/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CarPostgresRepository struct {
	db *pgxpool.Pool
}

func NewCarPostgresRepository(db *pgxpool.Pool) *CarPostgresRepository {
	return &CarPostgresRepository{db: db}
}

func (carRepository *CarPostgresRepository) GetAllCars() ([]models.Car, error) {
	rows, err := carRepository.db.Query(context.Background(), `
		SELECT id, brand, model, plate, year created_at
		FROM cars
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cars []models.Car

	for rows.Next() {
		var car models.Car

		err := rows.Scan(
			&car.ID,
			&car.Brand,
			&car.Model,
			&car.Plate,
			&car.Year,
			&car.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}

	return cars, nil
}

func (r *CarPostgresRepository) GetCarByPlate(plate string) (*models.Car, error) {

	query := `
		SELECT id, brand, model, plate, year, created_at
		FROM cars
		WHERE plate = $1
	`

	var car models.Car

	err := r.db.QueryRow(
		context.Background(),
		query,
		plate,
	).Scan(
		&car.ID,
		&car.Brand,
		&car.Model,
		&car.Plate,
		&car.Year,
		&car.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &car, nil
}

func (r *CarPostgresRepository) CreateCar(car models.Car) error {

	query := `
		INSERT INTO cars (id, brand, model, plate, year, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		car.ID,
		car.Brand,
		car.Model,
		car.Plate,
		car.Year,
		car.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
