package repositories

import "RentalAgency/models"

type CarRepository interface {
	ListAvailableCars() ([]models.Car, error)
	AddCar(car *models.Car) error
	RentCar(registration string) error
	ReturnCar(registration string, distanceDriven int) error
}
