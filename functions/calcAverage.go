package static

// CalcAverage calculates the average (mean) of a slice of integers.
func CalcAverage(data []int) int {
	// Initialize the variable to store the sum of all elements in the slice.
	average := 0

	// Iterate over each element in the data slice.
	for i := 0; i < len(data); i++ {
		// Add the current element's value to the average variable.
		average += data[i]
	}

	// Calculate and return the average by dividing the total sum by the number of elements.
	return average / len(data)
}
