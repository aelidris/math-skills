package static

// GetNumbers extracts integers from a string where each number is separated by new lines. It returns a slice of integers parsed from the string.
func GetNumbers(data string) []int {
	var result []int
	storeNumbers := ""

	for _, char := range data {

		// If the current character is a newline, convert the accumulated string (storeNumbers) to an integer
		if char == '\n' {
			if len(storeNumbers) > 0 {
				result = append(result, ConvertType(storeNumbers))
			}
			storeNumbers = ""
			continue
		}
		storeNumbers += string(char)
	}

	// Convert and store the last number
	if len(storeNumbers) != 0 {
		result = append(result, ConvertType(storeNumbers))
	}

	return result
}
