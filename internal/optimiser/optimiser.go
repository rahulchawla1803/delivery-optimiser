package optimiser

import (
	"fmt"
	"log"

	"github.com/rahulchawla1803/delivery-optimiser/internal/parser"
	"github.com/rahulchawla1803/delivery-optimiser/internal/timegraph"
)

// RouteStep represents a single step in the optimised route
type RouteStep struct {
	Type           string          `json:"type"` // "pickup" or "dropoff"
	OrderID        int             `json:"order_id"`
	RestaurantID   int             `json:"restaurant_id,omitempty"`
	RestaurantName string          `json:"restaurant_name,omitempty"`
	CustomerID     string          `json:"customer_id,omitempty"`
	CustomerName   string          `json:"customer_name,omitempty"`
	Location       parser.Location `json:"location"`
}

// Optimise processes the input, creates the time graph, and executes the selected algorithm
func Optimise(input parser.Input) ([]RouteStep, float64) {
	// Step 1: Create the time graph
	graph := timegraph.BuildTimeGraph(input)

	// Step 2: Select the optimization algorithm
	switch input.Config.Algorithm {
	case "greedy":
		return GreedyOptimise(input, graph)
	case "brute_force":
		return BruteForceOptimise(input, graph)
	case "dp":
		return DPOptimise(input, graph)
	default:
		log.Fatalf("Invalid algorithm choice: %s", input.Config.Algorithm)
	}

	return nil, 0
}

// BruteForceOptimise (Placeholder)
func BruteForceOptimise(input parser.Input, graph timegraph.TimeGraph) ([]RouteStep, float64) {
	fmt.Println("Executing Brute Force Algorithm...")
	// Implement brute force algorithm logic here
	return nil, 0
}

// DPOptimise (Placeholder)
func DPOptimise(input parser.Input, graph timegraph.TimeGraph) ([]RouteStep, float64) {
	fmt.Println("Executing DP Algorithm...")
	// Implement DP algorithm logic here
	return nil, 0
}
