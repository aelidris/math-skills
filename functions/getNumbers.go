package static

// GetNumbers extracts integers from a string where each number is separated by new lines.
// It returns a slice of integers parsed from the string.
func GetNumbers(data string) []int {
	var result []int   // Initialize a slice to store the parsed numbers
	storeNumbers := "" // Temporary string to accumulate characters of each number

	// Iterate over each character in the input string
	for _, char := range data {
		// If the current character is a newline, convert the accumulated string to an integer
		if char == '\n' {
			if len(storeNumbers) > 0 {
				result = append(result, ConvertType(storeNumbers)) // Convert and store the number
			}
			storeNumbers = "" // Reset the accumulator for the next number
			continue
		}
		// Add the character to the accumulator if it's not a newline
		storeNumbers += string(char)
	}

	// Handle any remaining characters that didn't end with a newline
	if len(storeNumbers) != 0 {
		result = append(result, ConvertType(storeNumbers)) // Convert and store the last number
	}

	return result // Return the slice of parsed integers
}
