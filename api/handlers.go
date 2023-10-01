package api

import (
	"RentalAgency/models"
	"RentalAgency/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var carService *services.CarService

func ListAvailableCars(w http.ResponseWriter, r *http.Request) {
	cars, err := carService.ListAvailableCars()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve cars")
		return
	}
	respondWithJSON(w, http.StatusOK, cars)
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := carService.AddCar(&car); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to add car")
		return
	}

	respondWithJSON(w, http.StatusCreated, car)
}

func RentCar(w http.ResponseWriter, r *http.Request) {
	registration := mux.Vars(r)["registration"]
	if err := carService.RentCar(registration); err != nil {
		respondWithError(w, http.StatusConflict, "Car is already rented")
		return
	}
	respondWithJSON(w, http.StatusOK, "Car rented successfully")
}

func ReturnCar(w http.ResponseWriter, r *http.Request) {
	registration := mux.Vars(r)["registration"]

	// Parse the request body for distance driven
	var request struct {
		DistanceDriven int `json:"distance_driven"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := carService.ReturnCar(registration, request.DistanceDriven); err != nil {
		respondWithError(w, http.StatusConflict, "Car is not rented or not found")
		return
	}
	respondWithJSON(w, http.StatusOK, "Car returned successfully")
}

func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	response := map[string]string{"error": message}
	respondWithJSON(w, status, response)
}
