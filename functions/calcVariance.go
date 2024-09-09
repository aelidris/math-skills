package static

import "math"

// CalcVariance calculates the variance of a slice of integers and returns a rounded integer.
func CalcVariance(data []int) int {
	// Calculate the mean (average)
	mean := CalcAverage(data)

	// Calculate the sum of squared differences from the mean
	var squaredDifferencesSum float64
	for _, num := range data {
		diff := float64(num - mean)
		squaredDifferencesSum += diff * diff
	}

	// Calculate the variance as the average of squared differences
	variance := squaredDifferencesSum / float64(len(data))

	// Round the variance to the nearest integer and convert to int
	return int(math.Round(variance))
}
