package static

// CalcVariance calculates the variance of a slice of integers
func CalcVariance(data []int) int {
	// Calculate the mean (average)
	mean := CalcAverage(data)

	// Calculate the squared differences from the mean
	var squaredDifferencesSum int
	for _, num := range data {
		diff := num - mean
		squaredDifferencesSum += diff * diff
	}

	// Calculate the variance
	variance := squaredDifferencesSum / len(data)
	return variance
}
