package static

import "strconv"

// ConvertType converts a string to an integer. If the conversion fails, it returns 0.
func ConvertType(s string) int {
	num := 0                    // Initialize a variable to store the converted integer
	num, err := strconv.Atoi(s) // Attempt to convert the string to an integer
	if err != nil {             // If there's an error during conversion
		return 0 // Return 0 as a default value
	}

	return num // Return the successfully converted integer
}
