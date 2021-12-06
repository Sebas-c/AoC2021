package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, errRead := fileRead("day6_input.txt")
	if errRead != nil {
		panic(errRead)
	}

	fmt.Printf("The amount of lanterfish after 80 days is: %d\n", fishCountTheSmartWay(input, 80))
	fmt.Printf("The amount of lanterfish after 256 days is: %d\n", fishCountTheSmartWay(input, 256))
	// fmt.Printf("The amount of lanterfish after 80 days is: %d\n", fishCount(input, 80))
	// fmt.Printf("The amount of lanterfish after 80 days is: %d\n", fishCount(input, 256)) can't handle :(
}

// Initial version, not efficient with large values for days. Also messes up the initial input.
func fishCount(currentArray []int, days int) int {
	tempArray := currentArray
	for i := 0; i < days; i++ {
		fmt.Printf("Processing day %d\n", i+1)
		tempArray = currentArray
		for i, fish := range tempArray {
			if fish == 0 {
				tempArray[i] = 6
				tempArray = append(tempArray, 8)
			} else {
				tempArray[i]--
			}
		}
		currentArray = tempArray
	}

	return len(tempArray)
}

func fishCountTheSmartWay(currentArray []int, days int) int {

	fishPerDay := make([]int, 9)
	totalFish := 0

	for i := range currentArray {
		fishPerDay[currentArray[i]]++
	}

	tempFishPerDay := fishPerDay
	for i := 0; i < days; i++ {
		tempFishPerDay = make([]int, 9)

		tempFishPerDay[8] = fishPerDay[0]
		tempFishPerDay[6] = fishPerDay[0]
		tempFishPerDay[0] = fishPerDay[1]
		tempFishPerDay[1] = fishPerDay[2]
		tempFishPerDay[2] = fishPerDay[3]
		tempFishPerDay[3] = fishPerDay[4]
		tempFishPerDay[4] = fishPerDay[5]
		tempFishPerDay[5] = fishPerDay[6]
		tempFishPerDay[6] += fishPerDay[7]
		tempFishPerDay[7] = fishPerDay[8]

		fishPerDay = tempFishPerDay
	}

	for i := range fishPerDay {
		totalFish += fishPerDay[i]
	}

	return totalFish
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
