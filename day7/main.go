package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, errRead := fileRead("day7_input.txt")
	if errRead != nil {
		panic(errRead)
	}

	fmt.Printf("Fuel used by the most cost-efficient solution (method 1): %d\n", findPosition(input))
	fmt.Printf("Fuel used by the most cost-efficient solution (method 2): %d\n", findPositionAgain(input))
}

func findPositionAgain(input []int) int {
	highest, _ := findExtremes(input)

	possiblePositions := make([]int, highest)

	for i := range possiblePositions {
		for j := range input {
			possiblePositions[i] += sumTo(int(math.Abs(float64(i - input[j]))))
		}
	}

	_, lowestFuel := findExtremes(possiblePositions)

	return lowestFuel
}

func sumTo(value int) int {
	total := 0
	for i := 0; i < value+1; i++ {
		total += i
	}
	return total
}

func findPosition(input []int) int {
	highest, _ := findExtremes(input)

	possiblePositions := make([]int, highest)

	for i := range possiblePositions {
		for j := range input {
			possiblePositions[i] += int(math.Abs(float64(i - input[j])))
		}
	}

	_, lowestFuel := findExtremes(possiblePositions)

	return lowestFuel
}

func findExtremes(input []int) (int, int) {
	highest := -1
	lowest := int(^uint(0) >> 1)

	for i, value := range input {
		if value > highest {
			highest = input[i]
		}

		if value < lowest {
			lowest = input[i]
		}
	}

	return highest, lowest
}

// Returns the file's content
func fileRead(filePath string) ([]int, error) {
	// Variables
	var input []int

	// Opening the file
	file, errFile := os.Open(filePath)

	if errFile != nil {
		return input, errFile
	}

	// Creating a scanner and setting it to read line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Appending every line
	scanner.Scan()
	line := scanner.Text()
	inputString := strings.Split(line, ",")

	for _, value := range inputString {
		valueInt, _ := strconv.Atoi(value)
		input = append(input, valueInt)
	}

	// Closing the file
	file.Close()

	return input, nil
}
