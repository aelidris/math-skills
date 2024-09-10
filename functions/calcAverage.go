package statistic

import "math"

// CalcAverage calculates the average of a slice of integers and returns a rounded integer.
func CalcAverage(data []int) int {
	sum := 0

	for i := 0; i < len(data); i++ {
		sum += data[i]
	}

	// Calculate the average as a float64 to ensure precision in rounding.
	average := float64(sum) / float64(len(data))

	return int(math.Round(average))
}
