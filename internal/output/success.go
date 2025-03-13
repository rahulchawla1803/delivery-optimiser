package output

import (
	"fmt"

	"github.com/rahulchawla1803/delivery-optimiser/internal/optimiser"
	"github.com/rahulchawla1803/delivery-optimiser/internal/utils"
)

// SuccessFlow handles successful output writing
func Success(result optimiser.OptimisedResult) error {
	// Step 1: Prepare Output Directory
	err := PrepareOutputDirectory()
	if err != nil {
		return fmt.Errorf("failed to prepare output directory: %v", err)
	}

	// Step 3: Generate text instructions
	instructions := generateTextInstructions(result)

	// Step 4: Write Output Files
	err = utils.WriteJSON("output/routes.json", result)
	if err != nil {
		return fmt.Errorf("failed to write routes.json: %v", err)
	}

	err = utils.WriteText("output/delivery_instructions.txt", instructions)
	if err != nil {
		return fmt.Errorf("failed to write delivery_instructions.txt: %v", err)
	}

	// Step 5: Write Debug Log
	debug := "Execution completed successfully."
	err = utils.WriteText("output/debug.log", debug)
	if err != nil {
		return fmt.Errorf("failed to write debug.log: %v", err)
	}

	return nil
}

// generateTextInstructions converts OptimisedResult to human-readable instructions
func generateTextInstructions(result optimiser.OptimisedResult) string {
	instructions := ""
	for _, step := range result.Route {
		if step.Type == "pickup" {
			instructions += fmt.Sprintf("- Pick up order #%d from %s\n", step.OrderID, step.RestaurantName)
		} else {
			instructions += fmt.Sprintf("- Deliver order #%d to Customer %s\n", step.OrderID, step.CustomerName)
		}
	}

	// Append Total Time at the end
	instructions += fmt.Sprintf("\nEstimated Total Time: %.2f minutes\n", result.TotalTime)

	return instructions
}
