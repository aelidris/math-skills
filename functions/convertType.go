package statistic

import (
	"log"
	"strconv"
)

// ConvertType converts a string to an integer
func ConvertType(s string) int {
	num := 0
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Conversion unsuccessful: failed to convert string to integer")
	}

	return num
}
