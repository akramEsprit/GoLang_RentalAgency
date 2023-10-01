package models

type Car struct {
	ID           int    `json:id`
	Model        string `json:"model"`
	Registration string `json:"registration"`
	Mileage      int    `json:"mileage"`
	IsRented     bool   `json:"is_rented"`
}
