package timegraph

import (
	"fmt"

	"github.com/rahulchawla1803/delivery-optimiser/internal/parser"
	"github.com/rahulchawla1803/delivery-optimiser/internal/utils"
)

// TimeGraph stores travel times between all points
type TimeGraph struct {
	Times map[string]map[string]float64 // Times[from][to] = travel time in minutes
}

// BuildTimeGraph computes travel times between all locations
func BuildTimeGraph(input parser.Input) TimeGraph {
	graph := TimeGraph{Times: make(map[string]map[string]float64)}

	// Initialize empty maps for each location
	allLocations := getAllLocations(input)
	for _, loc := range allLocations {
		graph.Times[loc] = make(map[string]float64)
	}

	// Compute travel times
	computeTimes(input, &graph)

	return graph
}

// getAllLocations generates unique location identifiers for driver, restaurants, and customers
func getAllLocations(input parser.Input) []string {
	var locations []string

	locations = append(locations, "driver")
	for _, r := range input.Restaurants {
		locations = append(locations, restaurantKey(r.ID))
	}
	for _, c := range input.Customers {
		locations = append(locations, customerKey(c.ID))
	}

	return locations
}

// computeTimes calculates travel time between all location pairs
func computeTimes(input parser.Input, graph *TimeGraph) {
	locations := getAllLocations(input)

	// Convert driver speed (km/h) to km/min
	speedKmph := float64(input.Driver.AvgSpeed)
	speedKmpm := speedKmph / 60.0

	// Compute all pairwise travel times
	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			locA, locB := locations[i], locations[j]

			latA, lonA := getLocationCoords(locA, input)
			latB, lonB := getLocationCoords(locB, input)

			distance := utils.HaversineDistance(latA, lonA, latB, lonB)
			time := distance / speedKmpm // Time in minutes

			// Store bidirectional times
			graph.Times[locA][locB] = time
			graph.Times[locB][locA] = time
		}
	}
}

// getLocationCoords returns the latitude & longitude for a given location key
func getLocationCoords(locKey string, input parser.Input) (float64, float64) {
	if locKey == "driver" {
		return input.Driver.Location.Latitude, input.Driver.Location.Longitude
	}

	for _, r := range input.Restaurants {
		if locKey == restaurantKey(r.ID) {
			return r.Location.Latitude, r.Location.Longitude
		}
	}

	for _, c := range input.Customers {
		if locKey == customerKey(c.ID) {
			return c.Location.Latitude, c.Location.Longitude
		}
	}

	return 0, 0 // Should never reach here
}

// Helper functions for location keys
func restaurantKey(id int) string  { return "R" + fmt.Sprint(id) }
func customerKey(id string) string { return "C" + id }
