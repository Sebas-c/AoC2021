package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, errRead := fileRead("day2_input.txt")
	if errRead != nil {
		panic(errRead)
	}

	fmt.Printf("The position is: %d\n", calculatePosition(input))
	fmt.Printf("The position with aim is: %d\n", calculatePositionWithAim(input))

}

func calculatePosition(input []string) int {
	positionX := 0
	positionY := 0

	for _, line := range input {
		lineSplit := strings.Split(line, " ")

		switch lineSplit[0] {
		case "up":
			positionY -= convertToInt(lineSplit[1])
		case "down":
			positionY += convertToInt(lineSplit[1])
		case "forward":
			positionX += convertToInt(lineSplit[1])
		default:
			panic("Invalid direction for input")
		}
	}
	return positionX * positionY
}

func calculatePositionWithAim(input []string) int {
	positionX := 0
	positionY := 0
	aim := 0

	for _, line := range input {
		lineSplit := strings.Split(line, " ")

		switch lineSplit[0] {
		case "up":
			aim -= convertToInt(lineSplit[1])
		case "down":
			aim += convertToInt(lineSplit[1])
		case "forward":
			positionX += convertToInt(lineSplit[1])
			positionY += aim * convertToInt(lineSplit[1])

		default:
			panic("Invalid direction for input")
		}
	}
	return positionX * positionY
}

func convertToInt(value string) int {
	valueInt, errConv := strconv.Atoi(value)
	if errConv != nil {
		panic(errConv)
	}

	return valueInt
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
