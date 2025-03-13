package optimiser

import (
	"math"

	"github.com/rahulchawla1803/delivery-optimiser/internal/parser"
	"github.com/rahulchawla1803/delivery-optimiser/internal/timegraph"
)

// BruteForceOptimise tries all possible routes and picks the best one
func BruteForceOptimise(input parser.Input, graph timegraph.TimeGraph) ([]RouteStep, float64) {
	var bestRoute []RouteStep
	minTime := math.MaxFloat64

	// Generate all possible order sequences (pickup + drop)
	allTasks := generateTasks(input)
	permutations := generatePermutations(allTasks)

	// Try all permutations to find the best one
	for _, perm := range permutations {
		route, totalTime := computeRouteTime(perm, graph)
		if totalTime < minTime {
			minTime = totalTime
			bestRoute = route
		}
	}

	return bestRoute, minTime
}

// generateTasks collects all pickup & drop tasks
func generateTasks(input parser.Input) []RouteStep {
	var tasks []RouteStep

	for _, order := range input.Orders {
		restaurant := getRestaurant(order.RestaurantID, input.Restaurants)
		customer := getCustomer(order.CustomerID, input.Customers)

		// Add pickup
		tasks = append(tasks, RouteStep{
			Type:           "pickup",
			OrderID:        order.OrderID,
			RestaurantID:   restaurant.ID,
			RestaurantName: restaurant.Name,
			Location:       restaurant.Location,
		})

		// Add dropoff
		tasks = append(tasks, RouteStep{
			Type:         "dropoff",
			OrderID:      order.OrderID,
			CustomerID:   customer.ID,
			CustomerName: customer.Name,
			Location:     customer.Location,
		})
	}

	return tasks
}

// generatePermutations returns all possible orderings of tasks
func generatePermutations(tasks []RouteStep) [][]RouteStep {
	var result [][]RouteStep
	permute(tasks, 0, &result)
	return result
}

// permute recursively generates all task permutations
func permute(tasks []RouteStep, start int, result *[][]RouteStep) {
	if start == len(tasks)-1 {
		permCopy := make([]RouteStep, len(tasks))
		copy(permCopy, tasks)
		*result = append(*result, permCopy)
		return
	}

	for i := start; i < len(tasks); i++ {
		tasks[start], tasks[i] = tasks[i], tasks[start] // Swap
		permute(tasks, start+1, result)
		tasks[start], tasks[i] = tasks[i], tasks[start] // Swap back
	}
}

// computeRouteTime calculates total travel time for a given task sequence
func computeRouteTime(tasks []RouteStep, graph timegraph.TimeGraph) ([]RouteStep, float64) {
	totalTime := 0.0
	currentLoc := "driver"

	for _, task := range tasks {
		locKey := getLocationKey(task)
		totalTime += graph.Times[currentLoc][locKey]
		currentLoc = locKey
	}

	return tasks, totalTime
}
