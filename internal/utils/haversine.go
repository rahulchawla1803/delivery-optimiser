package utils

import "math"

// EarthRadius is the approximate radius of Earth in kilometers
const EarthRadius = 6371.0

// HaversineDistance calculates the great-circle distance between two latitude-longitude points
func HaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	// Convert degrees to radians
	lat1, lon1, lat2, lon2 = degToRad(lat1), degToRad(lon1), degToRad(lat2), degToRad(lon2)

	// Apply Haversine formula
	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	return EarthRadius * c
}

// degToRad converts degrees to radians
func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
