package static

import "math"

// CalcStandardDeviation calculates the standard deviation of a slice of integers and returns a rounded integer.
func CalcStandardDeviation(data []int) int {
	// Calculate the variance
	variance := CalcVariance(data)

	// Calculate the standard deviation as the square root of the variance
	standardDeviation := math.Sqrt(float64(variance))

	// Round the standard deviation to the nearest integer and convert to int
	return int(math.Round(standardDeviation))
}
