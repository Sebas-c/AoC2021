package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board struct {
	values [][]int
	found  [][]bool
	won    bool
}

func main() {
	numbersDrawn, allBoards, errReading := fileRead("day4_input.txt")
	if errReading != nil {
		panic(errReading)
	}

	fmt.Printf("The score is: %d\n", findScore(numbersDrawn, allBoards))
	fmt.Printf("The score of the last winning board is: %d\n", findLastScore(numbersDrawn, allBoards))
}

func findScore(numbersDrawn []int, allBoards []board) int {
	for _, numberToDraw := range numbersDrawn {
		for i, board := range allBoards {
			markNumberAsFound(numberToDraw, board)

			if checkBingo(board) {
				allBoards[i].won = true
				sumAllUnmarked := sumAllUnmarked(board)

				return sumAllUnmarked * numberToDraw
			}
		}
	}

	return -1
}

func findLastScore(numbersDrawn []int, allBoards []board) int {
	boardswon := 1
	lastBoardFound := false

	for _, numberToDraw := range numbersDrawn {
		for i, board := range allBoards {
			if !board.won {
				markNumberAsFound(numberToDraw, board)

				if checkBingo(board) {
					allBoards[i].won = true
					boardswon += 1

					if lastBoardFound {
						return sumAllUnmarked(allBoards[i]) * numberToDraw
					}

					if boardswon == (len(allBoards) - 1) {
						lastBoardFound = true
					}
				}
			}
		}
	}

	return -1
}

func markNumberAsFound(numberDrawn int, currentBoard board) {
	for i, row := range currentBoard.values {
		for j, currentNumber := range row {
			if numberDrawn == currentNumber {

				currentBoard.found[i][j] = true
			}
		}
	}
}

func checkBingo(currentBoard board) bool {
	verticalBingo := false
	horizontalBingo := false

	for i := 0; i < len(currentBoard.found[0]); i++ {
		if currentBoard.found[i][0] && currentBoard.found[i][1] && currentBoard.found[i][2] && currentBoard.found[i][3] && currentBoard.found[i][4] {
			verticalBingo = true
		}
		if currentBoard.found[0][i] && currentBoard.found[1][i] && currentBoard.found[2][i] && currentBoard.found[3][i] && currentBoard.found[4][i] {
			horizontalBingo = true
		}
	}

	if horizontalBingo {
		return true
	}
	if verticalBingo {
		return true
	}

	return false
}

func sumAllUnmarked(currentBoard board) int {
	sum := 0

	for i, row := range currentBoard.found {
		for j, found := range row {
			if !found {
				sum += currentBoard.values[i][j]
			}
		}
	}
	return sum
}

func convertToInt(value string) (int, error) {
	valueInt, errConv := strconv.Atoi(value)
	if errConv != nil {
		return 0, errConv
	}

	return valueInt, nil
}

// Returns the file's content
func fileRead(filePath string) ([]int, []board, error) {
	// Variables
	var inputNumbers []int
	var inputBoards []board

	// Opening the file
	file, errFile := os.Open(filePath)

	if errFile != nil {
		return inputNumbers, inputBoards, errFile
	}

	// Creating a scanner and setting it to read line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Reading the numbers to be drawn

	Board := board{}
	boardLineCounter := 0
	firstLine := true

	// Appending every line
	for scanner.Scan() {
		line := scanner.Text()
		if firstLine {
			lineSplit := strings.Split(line, ",")
			for _, value := range lineSplit {
				valueInt, errConv := convertToInt(value)
				if errConv != nil {
					return inputNumbers, inputBoards, errConv
				}
				inputNumbers = append(inputNumbers, valueInt)
			}
			firstLine = false
		} else if line == "" {
			Board = board{}
			Board.values = make([][]int, 5)
			boardLineCounter = 0
		} else {
			valuesSplit := strings.Split(line, " ")
			for _, number := range valuesSplit {
				if number != "" {
					numberInt, errConv := convertToInt(number)
					if errConv != nil {
						return inputNumbers, inputBoards, errConv
					}
					Board.values[boardLineCounter] = append(Board.values[boardLineCounter], numberInt)
				}
			}
			boardLineCounter++
		}
		if boardLineCounter > 4 {
			boolMatrix := make([][]bool, 5)
			for index := range boolMatrix {
				boolMatrix[index] = make([]bool, 5)
			}
			Board.found = boolMatrix
			inputBoards = append(inputBoards, Board)
		}
	}

	// Closing the file
	file.Close()

	return inputNumbers, inputBoards, nil
}
