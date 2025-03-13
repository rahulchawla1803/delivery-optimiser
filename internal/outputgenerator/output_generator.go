package outputgenerator

import (
	"fmt"
	"log"

	"github.com/rahulchawla1803/delivery-optimiser/internal/optimiser"
	"github.com/rahulchawla1803/delivery-optimiser/internal/utils"
)

// GenerateOutput creates structured JSON and driver instructions
func GenerateOutput(route []optimiser.RouteStep, totalTime float64) error {
	// Convert route to structured output
	output := struct {
		Route            []optimiser.RouteStep `json:"route"`
		TotalTimeMinutes float64               `json:"total_time_minutes"`
	}{
		Route:            route,
		TotalTimeMinutes: totalTime,
	}

	// Generate driver instructions
	instructions := generateDriverInstructions(route, totalTime)

	// Write JSON output
	if err := utils.WriteJSON("output.json", output); err != nil {
		log.Fatalf("Failed to write output JSON: %v", err)
		return err
	}

	// Write text output for driver instructions
	if err := utils.WriteText("driver_instructions.txt", instructions); err != nil {
		log.Fatalf("Failed to write driver instructions: %v", err)
		return err
	}

	return nil
}

// generateDriverInstructions converts route into driver-friendly text format
func generateDriverInstructions(route []optimiser.RouteStep, totalTime float64) []string {
	var instructions []string
	instructions = append(instructions, "1. Start at your location.")

	stepNum := 2
	for _, step := range route {
		var instruction string
		if step.Type == "pickup" {
			instruction = fmt.Sprintf("%d. Pickup Order #%d from %s (%.4f, %.4f).",
				stepNum, step.OrderID, step.RestaurantName, step.Location.Latitude, step.Location.Longitude)
		} else if step.Type == "dropoff" {
			instruction = fmt.Sprintf("%d. Deliver Order #%d to %s (%.4f, %.4f).",
				stepNum, step.OrderID, step.CustomerName, step.Location.Latitude, step.Location.Longitude)
		}
		instructions = append(instructions, instruction)
		stepNum++
	}

	// Add final completion message
	instructions = append(instructions, fmt.Sprintf("%d. Route completed in approximately %.2f minutes.", stepNum, totalTime))

	return instructions
}
