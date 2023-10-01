// tests/handlers/car_handlers_test.go

package handlers

import (
	"RentalAgency/models"
	"RentalAgency/repositories"
	repository "RentalAgency/repositories"
	"RentalAgency/services"
	"testing"
)

func TestCarService_AddCar(t *testing.T) {
	// Create a mock repository for testing
	mockRepo := repositories.NewMockCarRepository()
	carService := services.NewCarService(mockRepo)

	// Define a car to add
	newCar := &models.Car{
		Model:        "Test Car",
		Registration: "TEST123",
		Mileage:      0,
		IsRented:     false,
	}

	// Add the car
	err := carService.AddCar(newCar)

	// Check for errors
	if err != nil {
		t.Errorf("AddCar returned an error: %v", err)
	}

	// Verify that the car was added to the mock repository
	cars, err := mockRepo.ListAvailableCars()
	if err != nil {
		t.Errorf("ListAllCars returned an error: %v", err)
	}

	// Ensure the car count is as expected
	if len(cars) != 2 {
		t.Errorf("Expected 2 cars, got %d", len(cars))
	}

	// Verify that the added car's details are correct
	addedCar := cars[len(cars)-1]
	if addedCar.Model != newCar.Model || addedCar.Registration != newCar.Registration || addedCar.Mileage != newCar.Mileage || addedCar.IsRented != newCar.IsRented {
		t.Errorf("Added car details do not match expected values")
	}
}

func TestCarService_ListAllCars(t *testing.T) {
	// Create a mock repository for testing
	mockRepo := repository.NewMockCarRepository()
	carService := services.NewCarService(mockRepo)

	// Add some test cars to the mock repository
	car1 := &models.Car{
		Model:        "Test Car 1",
		Registration: "TEST123",
		Mileage:      1000,
		IsRented:     false,
	}

	car2 := &models.Car{
		Model:        "Test Car 2",
		Registration: "TEST456",
		Mileage:      2000,
		IsRented:     true,
	}

	mockRepo.AddCar(car1)
	mockRepo.AddCar(car2)

	// List available cars
	allCars, err := carService.ListAvailableCars()
	if err != nil {
		t.Errorf("ListAllCars returned an error: %v", err)
	}

	// Ensure the correct number of available cars are returned
	if len(allCars) != 2 {
		t.Errorf("Expected 2 cars, got %d", len(allCars))
	}
}
