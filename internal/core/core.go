package core

import (
	"fmt"

	"github.com/rahulchawla1803/delivery-optimiser/internal/optimiser"
	"github.com/rahulchawla1803/delivery-optimiser/internal/output"
	"github.com/rahulchawla1803/delivery-optimiser/internal/parser"
	"github.com/rahulchawla1803/delivery-optimiser/internal/validation"
)

// Run executes the entire workflow
func Run() error {
	// Step 1: Load input.json
	input, err := parser.LoadInput("input.json")
	if err != nil {
		return output.Fail(fmt.Errorf("failed to load input.json: %v", err))
	}

	// Step 2: Load input_validation.json
	validationRules, err := parser.LoadValidation("input_validation.json")
	if err != nil {
		return output.Fail(fmt.Errorf("failed to load input_validation.json: %v", err))
	}

	// Step 3: Execute validation
	err = validation.ValidateInput(input, validationRules)
	if err != nil {
		return output.Fail(fmt.Errorf("validation failed: %v", err))
	}

	// Step 4: Run Optimisation
	// NOTE: Validation of rules shouldn't be in this scope.
	// The input should be validated beforehand, and the Haversine distance
	// should be computed and validated at an earlier stage (e.g., during parsing or validation).
	// Due to lack of time, this validation is being handled here.
	optimisedResult, err := optimiser.Optimise(input, validationRules)
	if err != nil {
		output.Fail(fmt.Errorf("optimiser failed: %v", err))
		return err
	}

	// Step 5: Handle Success Flow
	err = output.Success(optimisedResult)
	if err != nil {
		return output.Fail(fmt.Errorf("failed to process output: %v", err))
	}

	fmt.Println("Execution completed successfully.")
	return nil
}
