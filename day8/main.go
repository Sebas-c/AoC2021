package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, errRead := fileRead("day8_input.txt")
	if errRead != nil {
		panic(errRead)
	}
	fmt.Printf("Number of unique digits: %d\n", findUniqueDigits(input))
	fmt.Printf("Total of all outputs: %d\n", findTotalOutput(input))
}

func findTotalOutput(input []string) int {
	result := 0

	for _, line := range input {
		inputs := strings.Split(line, " ")
		patterns := decypherPatterns(inputs)
		for i := range patterns {
			patterns[i] = sortAlphabetically(patterns[i])
		}

		outputs := getOutputsAtLine(line)
		for i := range outputs {
			outputs[i] = sortAlphabetically(outputs[i])
		}

		currentLineValue := ""
		for _, output := range outputs {
			for i, pattern := range patterns {
				if strings.Compare(pattern, output) == 0 {
					currentLineValue += fmt.Sprint(i)
				}
			}
		}
		currentLineValueInt, _ := strconv.Atoi(currentLineValue)
		result += currentLineValueInt
	}

	return result
}

func sortAlphabetically(value string) string {
	sortedString := ""
	arrayValue := strings.Split(value, "")
	sort.Strings(arrayValue)
	for _, letter := range arrayValue {
		sortedString += letter
	}
	return sortedString
}

func getOutputsAtLine(line string) []string {
	lineSplit := strings.Split(line, "|")
	outputs := strings.Split(lineSplit[1], " ")

	return outputs
}

func decypherPatterns(input []string) []string {
	patterns := make([]string, 10)

	// Find patterns for "1", "4", "7" and "8" based on number of segments
	for i, pattern := range input {
		if len(pattern) == 2 {
			patterns[1] = input[i]
		}
		if len(pattern) == 3 {
			patterns[7] = input[i]
		}
		if len(pattern) == 4 {
			patterns[4] = input[i]
		}
		if len(pattern) == 7 {
			patterns[8] = input[i]
		}
	}

	// Find patterns for "2", "3" and "5" -- 5 segments numbers
	for i, pattern := range input {
		// If segments for "1" are both present, it's a "5"
		if len(pattern) == 5 {
			if strings.Contains(pattern, patterns[1][0:1]) && strings.Contains(pattern, patterns[1][1:2]) {
				patterns[3] = input[i]
			} else {
				// Count how many segments from "4" are present in the current pattern
				countMatches := 0
				for _, letter := range pattern {
					if strings.Compare(string(letter), patterns[4][0:1]) == 0 || strings.Compare(string(letter), patterns[4][1:2]) == 0 || strings.Compare(string(letter), patterns[4][2:3]) == 0 || strings.Compare(string(letter), patterns[4][3:4]) == 0 {
						countMatches++
					}
				}

				// If only 2 segments match, it's "2"
				if countMatches == 2 {
					patterns[2] = input[i]
					// If 3 segments match, it's "5"
				} else {
					patterns[5] = input[i]
				}
			}

			// Find patterns for "0", "6" and "9" -- 6 segments numbers
		} else if len(pattern) == 6 {
			// If all segments from a "4" are present, it's a "9"
			if strings.Contains(pattern, patterns[4][0:1]) && strings.Contains(pattern, patterns[4][1:2]) && strings.Contains(pattern, patterns[4][2:3]) && strings.Contains(pattern, patterns[4][3:4]) {
				patterns[9] = input[i]
				// It's not a "4" but contains all segments for "1", it's a "0"
			} else if strings.Contains(pattern, patterns[1][0:1]) && strings.Contains(pattern, patterns[1][1:2]) {
				patterns[0] = input[i]
				// It's not a "9" or "0" but has same number of segments, it's a "6"
			} else {
				patterns[6] = input[i]
			}
		}
	}

	return patterns
}

func findUniqueDigits(input []string) int {
	result := 0

	outputs := getOutputsOnly(input)

	for _, line := range outputs {
		for _, value := range line {
			if len(value) == 2 || len(value) == 3 || len(value) == 4 || len(value) == 7 {
				result += 1
			}
		}
	}

	return result
}

func getOutputsOnly(input []string) [][]string {
	outputs := make([][]string, len(input))
	for i, line := range input {
		// outputs[i] = make([]string, 4)

		lineSplit := strings.Split(line, "|")
		outputs[i] = strings.Split(lineSplit[1], " ")
	}

	return outputs
}

// Returns the file's content
func fileRead(filePath string) ([]string, error) {
	// Variables
	var input []string

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
		line := scanner.Text()
		input = append(input, line)
	}

	// Closing the file
	file.Close()

	return input, nil
}
