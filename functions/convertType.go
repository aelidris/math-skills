package static

import (
	"log"
	"strconv"
)

// ConvertType converts a string to an integer. If the conversion fails, it returns 0.
func ConvertType(s string) int {
	num := 0                    // Initialize a variable to store the converted integer
	num, err := strconv.Atoi(s) // Attempt to convert the string to an integer
	if err != nil {             // If there's an error during conversion
		log.Fatalf("Conversion unsuccessful: failed to convert string to integer")
	}

	return num // Return the successfully converted integer
}
