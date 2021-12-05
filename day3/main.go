package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	diagnostic, errRead := fileRead("day3_input.txt")
	if errRead != nil {
		panic(errRead)
	}

	fmt.Printf("The power consumption is: %d\n", powerConsumption(diagnostic))
	fmt.Printf("The life support rating is: %d\n", lifeSupportRating(diagnostic))
}

func lifeSupportRating(diagnostic []string) int {

	oxygen := findRating(diagnostic, 0, ">")
	co2 := findRating(diagnostic, 0, "<")

	return convertBinaryToDecimal(oxygen) * convertBinaryToDecimal(co2)

}

func findRating(tempArray []string, bitPosition int, comparator string) string {
	// Find most common bit at current position
	zeroes := 0
	ones := 0
	for _, value := range tempArray {
		if value[bitPosition:bitPosition+1] == "0" {
			zeroes++
		} else {
			ones++
		}
	}

	// Populate new array
	var newTempArray []string

	// Case zeroes and ones match
	if zeroes == ones {
		for _, row := range tempArray {
			if comparator == ">" {
				if row[bitPosition:bitPosition+1] == "1" {
					newTempArray = append(newTempArray, row)
				}
			} else {
				if row[bitPosition:bitPosition+1] == "0" {
					newTempArray = append(newTempArray, row)
				}
			}
		}
	}

	// More zeroes than ones
	if zeroes > ones {
		// Keeping the most common value
		if comparator == ">" {
			for _, row := range tempArray {
				if row[bitPosition:bitPosition+1] == "0" {
					newTempArray = append(newTempArray, row)
				}
			}
		} else { // Keeping the least common value
			for _, row := range tempArray {
				if row[bitPosition:bitPosition+1] == "1" {
					newTempArray = append(newTempArray, row)
				}
			}
		}
	}

	if ones > zeroes {
		// Keeping the most common value
		if comparator == ">" {
			for _, row := range tempArray {
				if row[bitPosition:bitPosition+1] == "1" {
					newTempArray = append(newTempArray, row)
				}
			}
		} else { // Keeping the least common value
			for _, row := range tempArray {
				if row[bitPosition:bitPosition+1] == "0" {
					newTempArray = append(newTempArray, row)
				}
			}
		}
	}

	if len(newTempArray) == 1 {
		return newTempArray[0]
	} else {
		return findRating(newTempArray, bitPosition+1, comparator)
	}
}

func powerConsumption(diagnostic []string) int {
	// Gamma is most common, epsilon is least common
	gamma := ""
	epsilon := ""

	for i := range diagnostic[0] {
		zeroes := 0
		ones := 0
		for _, bit := range diagnostic {
			if bit[i:i+1] == "0" {
				zeroes++
			} else {
				ones++
			}
		}

		if zeroes > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaDec, epsilonDec := convertBinaryToDecimal(gamma), convertBinaryToDecimal(epsilon)

	return gammaDec * epsilonDec
}

func convertBinaryToDecimal(value string) int {
	result, errConv := strconv.ParseInt(value, 2, 64)
	if errConv != nil {
		panic(errConv)
	}
	return int(result)
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
