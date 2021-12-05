package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, errRead := fileRead("day1_input.txt")
	if errRead != nil {
		panic(errRead)
	}

	fmt.Printf("Number of increases: %d", CountIncreases(input))
	fmt.Println()
	fmt.Printf("Number of window increases: %d", CountSlidingWindowIncreases(input))
	fmt.Println()

}

func CountIncreases(input []int) int {
	increasesCounter := 0

	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			increasesCounter += 1
		}
	}

	return increasesCounter
}

func CountSlidingWindowIncreases(input []int) int {
	windows := make([]int, len(input)-2)

	for index := range windows {
		windows[index] = input[index] + input[index+1] + input[index+2]
	}

	return CountIncreases(windows)
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
	for scanner.Scan() {
		line, errConversion := strconv.Atoi(scanner.Text())
		if errConversion != nil {
			return input, errConversion
		}
		input = append(input, line)
	}

	// Closing the file
	file.Close()

	return input, nil
}
