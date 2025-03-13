package core

import (
	"fmt"
	"log"

	"github.com/rahulchawla1803/delivery-optimiser/internal/optimiser"
	"github.com/rahulchawla1803/delivery-optimiser/internal/outputgenerator"
	"github.com/rahulchawla1803/delivery-optimiser/internal/parser"

	// "github.com/rahulchawla1803/delivery-optimiser/internal/utils"
	"github.com/rahulchawla1803/delivery-optimiser/internal/validation"
)

// Run loads JSON files and executes validation
func Run() error {
	// Load input.json
	input, err := parser.LoadInput("input.json")
	if err != nil {
		return fmt.Errorf("failed to load input.json: %v", err)
	}

	// Load input_validation.json
	validationRules, err := parser.LoadValidation("input_validation.json")
	if err != nil {
		return fmt.Errorf("failed to load input_validation.json: %v", err)
	}

	// PrintJSON
	// utils.PrintJSON("Validation Passed. Input is Valid.", input)
	// utils.PrintJSON("Parsed Validation Rules", validationRules)

	// Execute validation
	err = validation.ValidateInput(input, validationRules)
	if err != nil {
		return fmt.Errorf("validation failed: %v", err)
	}

	// Run Optimisation
	route, totalTime := optimiser.Optimise(input)

	// Generate Output (JSON + Driver Instructions)
	if err := outputgenerator.GenerateOutput(route, totalTime); err != nil {
		log.Fatalf("Error generating output: %v", err)
	}

	return nil
}
