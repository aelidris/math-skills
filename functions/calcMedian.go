package static

import (
	"math"
	"sort"
)

// CalcMedian calculates the median of a slice of integers and returns a rounded integer.
func CalcMedian(data []int) int {
	// Sort the slice of integers.
	sort.Ints(data)

	n := len(data)
	middle := n / 2

	if n%2 == 0 {
		// If even, median is the average of the two middle numbers.
		// Calculate the average as a float64 to ensure precision in rounding.
		median := float64(data[middle-1]+data[middle]) / 2
		// Round the average to the nearest integer and convert to int.
		return int(math.Round(median))
	} else {
		// If odd, median is the middle number.
		return data[middle]
	}
}
