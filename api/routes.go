package api

import (
	"RentalAgency/repositories"
	"RentalAgency/services"
	"RentalAgency/storage"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	carService = services.NewCarService(repositories.NewCarRepository(storage.DB))
	r.HandleFunc("/cars", ListAvailableCars).Methods("GET")
	r.HandleFunc("/cars", AddCar).Methods("POST")
	r.HandleFunc("/cars/{registration}/rentals", RentCar).Methods("POST")
	r.HandleFunc("/cars/{registration}/returns", ReturnCar).Methods("POST")
	return r
}
