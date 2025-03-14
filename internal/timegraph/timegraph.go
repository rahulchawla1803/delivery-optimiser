package timegraph

import (
	"fmt"

	"github.com/rahulchawla1803/delivery-optimiser/internal/parser"
	"github.com/rahulchawla1803/delivery-optimiser/internal/utils"
)

// TimeGraph stores travel times (in minutes) between all points.
type TimeGraph struct {
	Times map[string]map[string]float64 // Times[from][to] in minutes.
}

// BuildTimeGraph computes travel times between all locations and applies wait time for driver-to-restaurant legs.
func BuildTimeGraph(input parser.Input, rules parser.InputValidation) (TimeGraph, error) {
	graph := TimeGraph{Times: make(map[string]map[string]float64)}
	allLocations := getAllLocations(input)
	for _, loc := range allLocations {
		graph.Times[loc] = make(map[string]float64)
	}

	if err := computeTimes(input, &graph); err != nil {
		return graph, err
	}

	// (Optional) Validate each leg against allowed min/max if needed.
	// ...

	return graph, nil
}

// computeTimes calculates travel times between all location pairs.
// For driver->restaurant legs, it uses the maximum of the computed travel time and the restaurant's average wait time.
func computeTimes(input parser.Input, graph *TimeGraph) error {
	locations := getAllLocations(input)
	speedKmph := float64(input.Driver.AvgSpeed)
	speedKmpm := speedKmph / 60.0

	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			locA, locB := locations[i], locations[j]
			latA, lonA := getLocationCoords(locA, input)
			latB, lonB := getLocationCoords(locB, input)
			distance := utils.HaversineDistance(latA, lonA, latB, lonB)
			travelTime := distance / speedKmpm

			// If leg is from driver to a restaurant, use the maximum of travelTime and restaurant's avg wait time.
			if locA == "driver" && len(locB) > 0 && locB[0] == 'R' {
				r := findRestaurantByKey(locB, input)
				if float64(r.Preparation.AvgTime) > travelTime {
					travelTime = float64(r.Preparation.AvgTime)
				}
			}

			graph.Times[locA][locB] = travelTime
			graph.Times[locB][locA] = travelTime
		}
	}
	return nil
}

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
	return 0, 0
}

func restaurantKey(id int) string  { return "R" + fmt.Sprint(id) }
func customerKey(id string) string { return "C" + id }

// findRestaurantByKey retrieves a restaurant from its key.
func findRestaurantByKey(key string, input parser.Input) parser.Restaurant {
	var id int
	fmt.Sscanf(key, "R%d", &id)
	for _, r := range input.Restaurants {
		if r.ID == id {
			return r
		}
	}
	return parser.Restaurant{}
}
