package static

import "math"

// CalcStandardDeviation calculates the standard deviation of a slice of integers
func CalcStandardDeviation(data []int) int {
	// Calculate the variance
	variance := CalcVariance(data)

	// Calculate the standard deviation as the square root of the variance
	standardDeviation := math.Sqrt(float64(variance))

	// Return the standard deviation as an integer
	return int(standardDeviation)
}
