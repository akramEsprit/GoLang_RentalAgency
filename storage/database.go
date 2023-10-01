package storage

import (
	"RentalAgency/config"
	"RentalAgency/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := config.GetConfig()
	connectionString := config.DBConnectionString

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	DB.AutoMigrate(models.Car{})

	// Auto-migrate your database models here
}

func CloseDB() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}
