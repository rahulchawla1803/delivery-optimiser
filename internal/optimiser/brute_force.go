package optimiser

import (
	"math"

	"github.com/rahulchawla1803/delivery-optimiser/internal/parser"
	"github.com/rahulchawla1803/delivery-optimiser/internal/timegraph"
)

// BruteForceOptimise computes the optimal route by evaluating all valid permutations.
func BruteForceOptimise(input parser.Input, graph timegraph.TimeGraph) ([]RouteStep, float64) {
	// Precompute lookup maps.
	restaurantMap := make(map[int]parser.Restaurant)
	for _, r := range input.Restaurants {
		restaurantMap[r.ID] = r
	}
	customerMap := make(map[string]parser.Customer)
	for _, c := range input.Customers {
		customerMap[c.ID] = c
	}

	// Generate tasks: one pickup and one dropoff per order.
	var tasks []RouteStep
	for _, order := range input.Orders {
		r, ok := restaurantMap[order.RestaurantID]
		if !ok {
			continue
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

	// Generate all permutations.
	perms := generatePermutations(tasks)
	bestRoute := []RouteStep{}
	minTime := math.MaxFloat64

	// Evaluate each permutation.
	for _, perm := range perms {
		if !isValidPermutation(perm) {
			continue
		}
		totalTime := computeRouteTime(perm, graph, restaurantMap)
		if totalTime < minTime {
			minTime = totalTime
			bestRoute = perm
		}
	}

	return bestRoute, minTime
}

// isValidPermutation checks that for each order, pickup comes before dropoff.
func isValidPermutation(tasks []RouteStep) bool {
	pickupPos := make(map[int]int)
	dropoffPos := make(map[int]int)
	for i, task := range tasks {
		if task.Type == "pickup" {
			pickupPos[task.OrderID] = i
		} else {
			dropoffPos[task.OrderID] = i
		}
	}
	for orderID, pIdx := range pickupPos {
		if dIdx, ok := dropoffPos[orderID]; !ok || pIdx > dIdx {
			return false
		}
	}
	return true
}

// computeRouteTime calculates the total travel time for a given permutation of tasks,
// applying the wait time penalty for every pickup.
func computeRouteTime(tasks []RouteStep, graph timegraph.TimeGraph, restaurantMap map[int]parser.Restaurant) float64 {
	totalTime := 0.0
	currentLoc := "driver"
	for _, task := range tasks {
		key := getLocationKey(task)
		legTime := graph.Times[currentLoc][key]
		// For every pickup, apply the wait time penalty.
		if task.Type == "pickup" {
			r := restaurantMap[task.RestaurantID]
			waitTime := float64(r.Preparation.AvgTime)
			if waitTime > legTime {
				legTime = waitTime
			}
		}
		totalTime += legTime
		currentLoc = key
	}
	return totalTime
}

// generatePermutations recursively generates all permutations of tasks.
func generatePermutations(tasks []RouteStep) [][]RouteStep {
	var results [][]RouteStep
	permute(tasks, 0, &results)
	return results
}

func permute(tasks []RouteStep, start int, results *[][]RouteStep) {
	if start == len(tasks)-1 {
		permCopy := make([]RouteStep, len(tasks))
		copy(permCopy, tasks)
		*results = append(*results, permCopy)
		return
	}
	for i := start; i < len(tasks); i++ {
		tasks[start], tasks[i] = tasks[i], tasks[start]
		permute(tasks, start+1, results)
		tasks[start], tasks[i] = tasks[i], tasks[start] // backtrack
	}
}

// // getLocationKey returns a unique key for a task's location.
// func getLocationKey(task RouteStep) string {
// 	if task.Type == "pickup" {
// 		return "R" + fmt.Sprint(task.RestaurantID)
// 	}
// 	return "C" + task.CustomerID
// }
