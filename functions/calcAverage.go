package static

import "math"

// CalcAverage calculates the average (mean) of a slice of integers and returns a rounded integer.
func CalcAverage(data []int) int {
	// Initialize the variable to store the sum of all elements in the slice.
	sum := 0

	// Iterate over each element in the data slice.
	for i := 0; i < len(data); i++ {
		// Add the current element's value to the sum variable.
		sum += data[i]
	}

	// Calculate the average as a float64 to ensure precision in rounding.
	average := float64(sum) / float64(len(data))

	// Round the average to the nearest integer and convert to int.
	return int(math.Round(average))
}
