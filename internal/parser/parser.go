package parser

import (
	"github.com/rahulchawla1803/delivery-optimiser/internal/utils"
)

// Structs for JSON Parsing
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type DriverLimits struct {
	MaxDistance int `json:"max_distance"`
	MaxOrders   int `json:"max_orders"`
}

type Config struct {
	DriverLimits DriverLimits `json:"driver_limits"`
	Algorithm    string       `json:"algorithm"`
}

type Driver struct {
	Location Location `json:"location"`
	AvgSpeed int      `json:"avg_speed"`
}

type Order struct {
	OrderID      int    `json:"order_id"`
	RestaurantID int    `json:"restaurant_id"`
	CustomerID   string `json:"customer_id"`
}

type Restaurant struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Location    Location `json:"location"`
	Preparation struct {
		AvgTime    int     `json:"avg_time"`
		PeakFactor float64 `json:"peak_factor"`
	} `json:"preparation"`
}

type Customer struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

type Input struct {
	Config      Config       `json:"config"`
	Driver      Driver       `json:"driver"`
	Restaurants []Restaurant `json:"restaurants"`
	Customers   []Customer   `json:"customers"`
	Orders      []Order      `json:"orders"`
}

type InputValidation struct {
	Algorithm struct {
		AllowedValues []string `json:"allowed_values"`
	} `json:"algorithm"`
	DriverLimits struct {
		MaxDistance Range `json:"max_distance"`
		MaxOrders   Range `json:"max_orders"`
	} `json:"driver_limits"`
	Driver struct {
		AvgSpeed Range `json:"avg_speed"`
	} `json:"driver"`
	Restaurant struct {
		PrepTime   Range `json:"prep_time"`
		PeakFactor Range `json:"peak_factor"`
	} `json:"restaurant"`
	Location struct {
		Latitude  Range `json:"latitude"`
		Longitude Range `json:"longitude"`
	} `json:"location"`
}

type Range struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

// LoadInput loads input.json
func LoadInput(filePath string) (Input, error) {
	return utils.LoadJSON[Input](filePath)
}

// LoadValidation loads input_validation.json
func LoadValidation(filePath string) (InputValidation, error) {
	return utils.LoadJSON[InputValidation](filePath)
}
