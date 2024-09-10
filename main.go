package main

import (
	"fmt"
	"log"
	"os"

	statistic "statistic/functions"
)

func main() {
	// Check if there are enough arguments to start the program
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <datafile.txt>")
	}

	// Read the content of the file specified by the first argument
	fileContent, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to read the file %s: %v", os.Args[1], err)
	}

	// Check if the file is empty
	if len(fileContent) == 0 {
		log.Fatalf("There is no data in <%s>", os.Args[1])
	}

	// Convert the file content to a slice of integers
	data := statistic.GetNumbers(string(fileContent))

	// Calculate and print the average of the numbers
	fmt.Printf("Average: %d\n", statistic.CalcAverage(data))

	// Calculate and print the median of the numbers
	fmt.Printf("Median: %d\n", statistic.CalcMedian(data))

	// Calculate and print the variance of the numbers
	fmt.Printf("Variance: %d\n", statistic.CalcVariance(data))

	// Calculate and print the standard deviation of the numbers
	fmt.Printf("Standard Deviation: %d\n", statistic.CalcStandardDeviation(data))
}
