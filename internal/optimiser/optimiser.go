package optimiser

import (
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

// OptimisedResult wraps the final route and total time
type OptimisedResult struct {
	Route     []RouteStep `json:"route"`
	TotalTime float64     `json:"total_time"`
}

// Optimise processes the input, creates the time graph, and executes the selected algorithm
func Optimise(input parser.Input) OptimisedResult {
	// Step 1: Create the time graph
	graph := timegraph.BuildTimeGraph(input)

	// Step 2: Select the optimization algorithm
	var route []RouteStep
	var totalTime float64

	switch input.Config.Algorithm {
	case "greedy":
		route, totalTime = GreedyOptimise(input, graph)
	case "brute_force":
		route, totalTime = BruteForceOptimise(input, graph)
	case "dp":
		// route, totalTime = DPOptimise(input, graph)
	default:
		log.Fatalf("Invalid algorithm choice: %s", input.Config.Algorithm)
	}

	return OptimisedResult{
		Route:     route,
		TotalTime: totalTime,
	}
}
