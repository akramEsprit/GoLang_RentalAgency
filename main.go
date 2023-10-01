package main

import (
	"RentalAgency/api"
	"RentalAgency/config"
	"RentalAgency/storage"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := config.LoadConfig()
	storage.InitDB()
	defer storage.CloseDB()

	r := api.SetupRoutes()
	serverPort := config.ServerPort
	fmt.Println("Server is running on port " + serverPort)
	log.Fatal(http.ListenAndServe(serverPort, r))
}
