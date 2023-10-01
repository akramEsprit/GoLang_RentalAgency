package services

import (
	"RentalAgency/models"
	repository "RentalAgency/repositories"
)

type CarService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) *CarService {
	return &CarService{repo}
}

func (s *CarService) ListAvailableCars() ([]models.Car, error) {
	return s.repo.ListAvailableCars()
}

func (s *CarService) AddCar(car *models.Car) error {
	return s.repo.AddCar(car)
}

func (s *CarService) RentCar(registration string) error {
	return s.repo.RentCar(registration)
}

func (s *CarService) ReturnCar(registration string, distanceDriven int) error {
	return s.repo.ReturnCar(registration, distanceDriven)
}
