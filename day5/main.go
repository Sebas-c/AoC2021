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
	input, errRead := fileRead("day5_input.txt")
	if errRead != nil {
		panic(errRead)
	}

	fmt.Printf("Number of points where two lines or more overlap: %d\n", countDoubleOverlap(populateMatrix(input, 999, false)))
	fmt.Printf("Number of points where two lines or more overlap (including diagonals): %d\n", countDoubleOverlap(populateMatrix(input, 999, true)))

}

func countDoubleOverlap(matrix [][]int) int {
	counter := 0

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] >= 2 {
				counter++
			}
		}
	}

	return counter
}

func populateMatrix(input []string, matrixSize int, useDiagonals bool) [][]int {
	// Create matrix and init all values to 0
	matrix := make([][]int, matrixSize)
	for i := range matrix {
		matrix[i] = make([]int, matrixSize)
	}

	// Breakdown input
	for _, value := range input {
		valuesplit := strings.Split(value, " -> ")
		startPositionSplit := strings.Split(valuesplit[0], ",")
		endPositionSplit := strings.Split(valuesplit[1], ",")

		startPositionInt := []int{convertToInt(startPositionSplit[0]), convertToInt(startPositionSplit[1])}
		endPositionInt := []int{convertToInt(endPositionSplit[0]), convertToInt(endPositionSplit[1])}

		// Matching X position
		if startPositionInt[0] == endPositionInt[0] {
			if startPositionInt[1] < endPositionInt[1] {
				for i := startPositionInt[1]; i < endPositionInt[1]+1; i++ {
					matrix[startPositionInt[0]][i]++
				}
			} else {
				for i := endPositionInt[1]; i < startPositionInt[1]+1; i++ {
					matrix[startPositionInt[0]][i]++
				}
			}
		}
		// Matchin Y position
		if startPositionInt[1] == endPositionInt[1] {
			if startPositionInt[0] < endPositionInt[0] {
				for i := startPositionInt[0]; i < endPositionInt[0]+1; i++ {
					matrix[i][startPositionInt[1]]++
				}
			} else {
				for i := endPositionInt[0]; i < startPositionInt[0]+1; i++ {
					matrix[i][startPositionInt[1]]++
				}
			}
		}

		// Checking diagonals
		if useDiagonals {
			if math.Abs(float64(startPositionInt[0]-endPositionInt[0])) == math.Abs(float64(startPositionInt[1]-endPositionInt[1])) {
				posX, posY := startPositionInt[0], startPositionInt[1]

				// startX < endX && startY < endY -> X++ && Y++
				if startPositionInt[0] < endPositionInt[0] && startPositionInt[1] < endPositionInt[1] {
					for i := startPositionInt[0]; i < endPositionInt[0]+1; i++ {
						matrix[posX][posY]++
						posX++
						posY++
					}
				}

				// startX < endX && startY > endY -> X++ && Y--
				if startPositionInt[0] < endPositionInt[0] && startPositionInt[1] > endPositionInt[1] {
					for i := startPositionInt[0]; i < endPositionInt[0]+1; i++ {
						matrix[posX][posY]++
						posX++
						posY--
					}
				}

				// startX > endX && startY < endY -> X-- && Y++
				if startPositionInt[0] > endPositionInt[0] && startPositionInt[1] < endPositionInt[1] {
					for i := endPositionInt[0]; i < startPositionInt[0]+1; i++ {
						matrix[posX][posY]++
						posX--
						posY++
					}
				}

				// startX > endX && startY > endY -> X-- && Y--
				if startPositionInt[0] > endPositionInt[0] && startPositionInt[1] > endPositionInt[1] {
					for i := endPositionInt[0]; i < startPositionInt[0]+1; i++ {
						matrix[posX][posY]++
						posX--
						posY--
					}
				}
			}
		}
	}

	return matrix
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
