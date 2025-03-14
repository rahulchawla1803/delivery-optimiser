package optimiser

import (
	"fmt"
	"sort"

	"github.com/rahulchawla1803/delivery-optimiser/internal/parser"
	"github.com/rahulchawla1803/delivery-optimiser/internal/timegraph"
)

// GreedyOptimise selects the next task based on effective cost.
// For pickups, effective cost = max(travel time from current location, restaurant's avg wait time).
func GreedyOptimise(input parser.Input, graph timegraph.TimeGraph) ([]RouteStep, float64) {
	var route []RouteStep
	currentLoc := "driver"
	totalTime := 0.0
	pickedUp := make(map[int]bool)

	// Precompute lookup maps for restaurants and customers.
	restaurantMap := make(map[int]parser.Restaurant)
	for _, r := range input.Restaurants {
		restaurantMap[r.ID] = r
	}
	customerMap := make(map[string]parser.Customer)
	for _, c := range input.Customers {
		customerMap[c.ID] = c
	}

	// Create tasks: one pickup and one dropoff per order.
	var tasks []RouteStep
	for _, order := range input.Orders {
		r, ok := restaurantMap[order.RestaurantID]
		if !ok {
			continue // or handle error
		}
		c, ok := customerMap[order.CustomerID]
		if !ok {
			continue
		}
		tasks = append(tasks, RouteStep{
			Type:           "pickup",
			OrderID:        order.OrderID,
			RestaurantID:   r.ID,
			RestaurantName: r.Name,
			Location:       r.Location,
		})
		tasks = append(tasks, RouteStep{
			Type:         "dropoff",
			OrderID:      order.OrderID,
			CustomerID:   c.ID,
			CustomerName: c.Name,
			Location:     c.Location,
		})
	}

	// Process tasks until none remain.
	for len(tasks) > 0 {
		// Filter valid tasks: pickups are always valid; dropoffs only if pickup has been done.
		var validTasks []RouteStep
		for _, task := range tasks {
			if task.Type == "pickup" || (task.Type == "dropoff" && pickedUp[task.OrderID]) {
				validTasks = append(validTasks, task)
			}
		}
		if len(validTasks) == 0 {
			break
		}

		// Sort valid tasks by effective cost from current location.
		sort.Slice(validTasks, func(i, j int) bool {
			costI := graph.Times[currentLoc][getLocationKey(validTasks[i])]
			costJ := graph.Times[currentLoc][getLocationKey(validTasks[j])]
			// For pickups, add wait time penalty in cost computation.
			if validTasks[i].Type == "pickup" {
				r := restaurantMap[validTasks[i].RestaurantID]
				costI = max(costI, float64(r.Preparation.AvgTime))
			}
			if validTasks[j].Type == "pickup" {
				r := restaurantMap[validTasks[j].RestaurantID]
				costJ = max(costJ, float64(r.Preparation.AvgTime))
			}
			return costI < costJ
		})

		nextTask := validTasks[0]
		travelTime := graph.Times[currentLoc][getLocationKey(nextTask)]
		// If it's a pickup, effective time is max(travel time, wait time)
		if nextTask.Type == "pickup" {
			r := restaurantMap[nextTask.RestaurantID]
			travelTime = max(travelTime, float64(r.Preparation.AvgTime))
		}
		totalTime += travelTime
		currentLoc = getLocationKey(nextTask)
		route = append(route, nextTask)
		tasks = removeTask(tasks, nextTask)
		if nextTask.Type == "pickup" {
			pickedUp[nextTask.OrderID] = true
		}
	}

	return route, totalTime
}

// removeTask removes a specific task from the slice.
func removeTask(tasks []RouteStep, taskToRemove RouteStep) []RouteStep {
	for i, task := range tasks {
		if task == taskToRemove {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

// getLocationKey returns a unique key for a task's location.
func getLocationKey(task RouteStep) string {
	if task.Type == "pickup" {
		return "R" + fmt.Sprint(task.RestaurantID)
	}
	return "C" + task.CustomerID
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
