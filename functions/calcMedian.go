package static

import "sort"

// CalcMedian calculates the median of a slice of integers
func CalcMedian(data []int) int {
	// Sort the slice of integers
	sort.Ints(data)

	n := len(data)
	middle := n / 2

	if n%2 == 0 {
		// If even, median is the average of the two middle numbers
		return (data[middle-1] + data[middle]) / 2
	} else {
		// If odd, median is the middle number
		return data[middle]
	}
}
