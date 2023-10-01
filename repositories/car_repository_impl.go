package repositories

import (
	"RentalAgency/models"

	"gorm.io/gorm"
)

type CarRepositoryImpl struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepositoryImpl {
	return &CarRepositoryImpl{db}
}

func (r *CarRepositoryImpl) ListAvailableCars() ([]models.Car, error) {
	var cars []models.Car
	if err := r.db.Where("is_rented = ?", false).Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func (r *CarRepositoryImpl) AddCar(car *models.Car) error {
	if err := r.db.Create(car).Error; err != nil {
		return err
	}
	return nil
}

func (r *CarRepositoryImpl) RentCar(registration string) error {
	var car models.Car
	if err := r.db.First(&car, "registration = ?", registration).Error; err != nil {
		return err
	}

	if car.IsRented {
		return nil // Car is already rented
	}

	car.IsRented = true
	if err := r.db.Save(&car).Error; err != nil {
		return err
	}

	return nil
}

func (r *CarRepositoryImpl) ReturnCar(registration string, distanceDriven int) error {
	var car models.Car
	if err := r.db.First(&car, "registration = ?", registration).Error; err != nil {
		return err
	}

	if !car.IsRented {
		return nil // Car is not rented
	}

	car.Mileage += distanceDriven
	car.IsRented = false
	if err := r.db.Save(&car).Error; err != nil {
		return err
	}

	return nil
}
