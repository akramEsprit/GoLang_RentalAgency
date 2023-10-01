package repositories

import (
	"RentalAgency/models"
	"errors"
)

/*
*
MockCarRepository is a mock implementation of CarRepository for testing purposes.
*/
type MockCarRepository struct {
	Cars []models.Car
}

// NewMockCarRepository creates a new instance of MockCarRepository.
func NewMockCarRepository() *MockCarRepository {
	return &MockCarRepository{
		Cars: []models.Car{
			{
				ID:           1,
				Model:        "Car 1",
				Registration: "Registration1",
				Mileage:      1000,
				IsRented:     false,
			},
			{
				ID:           2,
				Model:        "Car 2",
				Registration: "Registration2",
				Mileage:      2000,
				IsRented:     true,
			},
		},
	}
}

// ListAvailableCars returns a list of available cars from the mock repository.
func (r *MockCarRepository) ListAvailableCars() ([]models.Car, error) {
	var availableCars []models.Car
	for _, car := range r.Cars {
		if !car.IsRented {
			availableCars = append(availableCars, car)
		}
	}
	return availableCars, nil
}

// AddCar adds a car to the mock repository.
func (r *MockCarRepository) AddCar(car *models.Car) error {
	car.ID = len(r.Cars) + 1
	r.Cars = append(r.Cars, *car)
	return nil
}

// RentCar marks a car as rented in the mock repository.
func (r *MockCarRepository) RentCar(registration string) error {
	for i, car := range r.Cars {
		if car.Registration == registration {
			r.Cars[i].IsRented = true
			return nil
		}
	}
	return errors.New("car not found")
}

// ReturnCar marks a car as available in the mock repository.
func (r *MockCarRepository) ReturnCar(registration string, distanceDriven int) error {
	for i, car := range r.Cars {
		if car.Registration == registration {
			r.Cars[i].IsRented = false
			r.Cars[i].Mileage += distanceDriven
			return nil
		}
	}
	return errors.New("car not found")
}
