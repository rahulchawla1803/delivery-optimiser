package optimiser

import (
	"fmt"
	"sort"

	"github.com/rahulchawla1803/delivery-optimiser/internal/parser"
	"github.com/rahulchawla1803/delivery-optimiser/internal/timegraph"
)

// GreedyOptimise selects the nearest available location iteratively
func GreedyOptimise(input parser.Input, graph timegraph.TimeGraph) ([]RouteStep, float64) {
	var route []RouteStep
	visited := make(map[string]bool)
	currentLoc := "driver"
	totalTime := 0.0

	// Collect all pickup and drop-off points
	var tasks []RouteStep
	for _, order := range input.Orders {
		restaurant := getRestaurant(order.RestaurantID, input.Restaurants)
		customer := getCustomer(order.CustomerID, input.Customers)

		// Add pickup tasks
		tasks = append(tasks, RouteStep{
			Type:           "pickup",
			OrderID:        order.OrderID,
			RestaurantID:   restaurant.ID,
			RestaurantName: restaurant.Name,
			Location:       restaurant.Location,
		})

		// Add drop-off tasks
		tasks = append(tasks, RouteStep{
			Type:         "dropoff",
			OrderID:      order.OrderID,
			CustomerID:   customer.ID,
			CustomerName: customer.Name,
			Location:     customer.Location,
		})
	}

	// Process all tasks using a greedy nearest-neighbor approach
	for len(tasks) > 0 {
		// Sort tasks based on shortest travel time from the current location
		sort.Slice(tasks, func(i, j int) bool {
			return graph.Times[currentLoc][getLocationKey(tasks[i])] < graph.Times[currentLoc][getLocationKey(tasks[j])]
		})

		// Pick the nearest task
		nextTask := tasks[0]
		tasks = tasks[1:]

		// Update total time & mark visited
		totalTime += graph.Times[currentLoc][getLocationKey(nextTask)]
		currentLoc = getLocationKey(nextTask)
		visited[currentLoc] = true

		// Append to final route
		route = append(route, nextTask)
	}

	return route, totalTime
}

// Helper to get restaurant details
func getRestaurant(id int, restaurants []parser.Restaurant) parser.Restaurant {
	for _, r := range restaurants {
		if r.ID == id {
			return r
		}
	}
	return parser.Restaurant{}
}

// Helper to get customer details
func getCustomer(id string, customers []parser.Customer) parser.Customer {
	for _, c := range customers {
		if c.ID == id {
			return c
		}
	}
	return parser.Customer{}
}

// Helper to get a unique location key
func getLocationKey(step RouteStep) string {
	if step.Type == "pickup" {
		return "R" + fmt.Sprint(step.RestaurantID)
	}
	return "C" + step.CustomerID
}
