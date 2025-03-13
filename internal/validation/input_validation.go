package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/rahulchawla1803/delivery-optimiser/internal/parser"
)

// Validator instance
var validate = validator.New()

// ValidateInput performs field validation and custom validation using input_validation.json
func ValidateInput(input parser.Input, rules parser.InputValidation) error {
	// Step 1: Validate required fields using go-playground/validator
	if err := validate.Struct(input); err != nil {
		return fmt.Errorf("field validation failed: %v", err)
	}

	// Step 2: Validate numeric constraints using rules from input_validation.json
	if err := validateAlgorithm(input.Config, rules); err != nil {
		return err
	}
	if err := validateDriverLimits(input, rules); err != nil {
		return err
	}
	if err := validateDriver(input.Driver, rules); err != nil {
		return err
	}
	if err := validateRestaurants(input.Restaurants, rules); err != nil {
		return err
	}
	if err := validateCustomers(input.Customers); err != nil {
		return err
	}
	if err := validateOrders(input.Orders, input.Restaurants, input.Customers); err != nil {
		return err
	}

	return nil
}

// Validate algorithm
func validateAlgorithm(config parser.Config, rules parser.InputValidation) error {
	allowed := rules.Algorithm.AllowedValues
	for _, valid := range allowed {
		if config.Algorithm == valid {
			return nil
		}
	}
	return fmt.Errorf("invalid algorithm selection: %s. Allowed values: %v", config.Algorithm, allowed)
}

// Validate driver limits
func validateDriverLimits(input parser.Input, rules parser.InputValidation) error {
	if input.Config.DriverLimits.MaxDistance < int(rules.DriverLimits.MaxDistance.Min) ||
		input.Config.DriverLimits.MaxDistance > int(rules.DriverLimits.MaxDistance.Max) {
		return fmt.Errorf("max_distance must be between %v and %v", rules.DriverLimits.MaxDistance.Min, rules.DriverLimits.MaxDistance.Max)
	}

	if input.Config.DriverLimits.MaxOrders < int(rules.DriverLimits.MaxOrders.Min) ||
		input.Config.DriverLimits.MaxOrders > int(rules.DriverLimits.MaxOrders.Max) {
		return fmt.Errorf("max_orders must be between %v and %v", rules.DriverLimits.MaxOrders.Min, rules.DriverLimits.MaxOrders.Max)
	}
	return nil
}

// Validate driver attributes
func validateDriver(driver parser.Driver, rules parser.InputValidation) error {
	if driver.AvgSpeed < int(rules.Driver.AvgSpeed.Min) || driver.AvgSpeed > int(rules.Driver.AvgSpeed.Max) {
		return fmt.Errorf("driver avg_speed must be between %v and %v", rules.Driver.AvgSpeed.Min, rules.Driver.AvgSpeed.Max)
	}
	if err := validateLocation(driver.Location, rules); err != nil {
		return fmt.Errorf("invalid driver location: %v", err)
	}
	return nil
}

// Validate restaurant attributes
func validateRestaurants(restaurants []parser.Restaurant, rules parser.InputValidation) error {
	restaurantIDs := make(map[int]bool)

	for _, restaurant := range restaurants {
		// Ensure unique restaurant ID
		if restaurantIDs[restaurant.ID] {
			return fmt.Errorf("duplicate restaurant_id found: %d", restaurant.ID)
		}
		restaurantIDs[restaurant.ID] = true

		// Validate location
		if err := validateLocation(restaurant.Location, rules); err != nil {
			return fmt.Errorf("invalid restaurant location for id %d: %v", restaurant.ID, err)
		}

		// Validate preparation time
		if restaurant.Preparation.AvgTime < int(rules.Restaurant.PrepTime.Min) ||
			restaurant.Preparation.AvgTime > int(rules.Restaurant.PrepTime.Max) {
			return fmt.Errorf("prep_time for restaurant %d must be between %v and %v",
				restaurant.ID, rules.Restaurant.PrepTime.Min, rules.Restaurant.PrepTime.Max)
		}

		// Validate peak factor
		if restaurant.Preparation.PeakFactor < rules.Restaurant.PeakFactor.Min ||
			restaurant.Preparation.PeakFactor > rules.Restaurant.PeakFactor.Max {
			return fmt.Errorf("peak_factor for restaurant %d must be between %v and %v",
				restaurant.ID, rules.Restaurant.PeakFactor.Min, rules.Restaurant.PeakFactor.Max)
		}
	}

	return nil
}

// Validate customers
func validateCustomers(customers []parser.Customer) error {
	customerIDs := make(map[string]bool)

	for _, customer := range customers {
		// Ensure unique customer ID
		if customerIDs[customer.ID] {
			return fmt.Errorf("duplicate customer_id found: %s", customer.ID)
		}
		customerIDs[customer.ID] = true

		// Validate customer location
		if customer.Location.Latitude == 0 || customer.Location.Longitude == 0 {
			return fmt.Errorf("invalid location for customer_id: %s", customer.ID)
		}
	}

	return nil
}

// Validate orders
func validateOrders(orders []parser.Order, restaurants []parser.Restaurant, customers []parser.Customer) error {
	orderIDs := make(map[int]bool)
	restaurantSet := make(map[int]bool)
	customerSet := make(map[string]bool)

	// Create lookup for restaurant and customer IDs
	for _, r := range restaurants {
		restaurantSet[r.ID] = true
	}
	for _, c := range customers {
		customerSet[c.ID] = true
	}

	for _, order := range orders {
		// Ensure unique order ID
		if orderIDs[order.OrderID] {
			return fmt.Errorf("duplicate order_id found: %d", order.OrderID)
		}
		orderIDs[order.OrderID] = true

		// Validate restaurant exists
		if !restaurantSet[order.RestaurantID] {
			return fmt.Errorf("order_id %d references invalid restaurant_id %d", order.OrderID, order.RestaurantID)
		}

		// Validate customer exists
		if !customerSet[order.CustomerID] {
			return fmt.Errorf("order_id %d references invalid customer_id %s", order.OrderID, order.CustomerID)
		}
	}

	return nil
}

// Validate latitude and longitude
func validateLocation(location parser.Location, rules parser.InputValidation) error {
	if location.Latitude < rules.Location.Latitude.Min || location.Latitude > rules.Location.Latitude.Max {
		return errors.New("latitude out of range")
	}
	if location.Longitude < rules.Location.Longitude.Min || location.Longitude > rules.Location.Longitude.Max {
		return errors.New("longitude out of range")
	}
	return nil
}
