# Statistical Calculator

This Go-based application calculates various statistical metrics—such as average, median, variance, and standard deviation—from numerical data provided in a text file. The application reads the data from the file, processes it, and outputs the results to the console.

## Description

The Statistical Calculator performs the following operations:
- Calculate Average: Computes the arithmetic mean of the dataset.
- Calculate Median: Determines the middle value of the dataset.
- Calculate Variance: Measures the spread of the numbers in the dataset.
- Calculate Standard Deviation: Evaluates the amount of variation or dispersion in the dataset.

## Author
- [ABDELFATTAH ELIDRISSI]

## Usage

### Prerequisites

- Go (version 1.22.3 or later) installed on your system.

### How to Run

1. **Clone the Repository:**

```sh
git clone https://learn.zone01oujda.ma/git/aelidris/math-skills
cd math-skills
```
2. **Compile and Run:** To compile and run the application, use the following command:

```sh
go run main.go data.txt

```

Replace data.txt with the path to your text file containing the numerical data.

### Input File Format

- The input file should be a plain text file with one number per line.
Example:

```sh
10
20
30
40
50
```

### Output

The application will output the calculated average, median, variance, and standard deviation to the console.

### Example Output

```sh
Average: 30
Median: 30
Variance: 200
Standard Deviation: 14
```

## Implementation Details

### Project Structure

Main Go File (main.go): Contains the entry point of the application and orchestrates the reading of data, processing, and displaying results.
Static Package (static): Contains utility functions such as GetNumbers, CalcAverage, CalcMedian, CalcVariance, and CalcStandardDeviation.

### Functions Overview

- ConvertType(string) int: Converts a string to an integer, returning 0 if the conversion fails.
- GetNumbers(string) []int: Parses the content of the input file and returns a slice of integers.
- CalcAverage([]int) int: Calculates and returns the average of the input data.
- CalcMedian([]int) int: Calculates and returns the median of the input data.
- CalcVariance([]int) int: Calculates and returns the variance of the input data.
- CalcStandardDeviation([]int) int: Calculates and returns the standard deviation of the input data.

### Error Handling

The application includes appropriate error handling:
- Usage Error: If no input file is provided, the application will print a usage message and terminate.
- File Reading Error: If the input file cannot be read, the application will print an error message and terminate.
- Empty File Error: If the input file is empty, the application will print an error message and terminate.