package main

import "testing"

func TestCheckBingoPositiveVertical(t *testing.T) {
	testBoard := board{
		found: [][]bool{
			{true, false, false, false, false},
			{true, false, false, false, false},
			{true, false, false, false, false},
			{true, false, false, false, false},
			{true, false, false, false, false}},
	}

	if !checkBingo(testBoard) {
		t.Error("Whoopsie daisy, func checkBingo is broken (vertical bingo)")
	}
}

func TestCheckBingoPositiveHorizontal(t *testing.T) {
	testBoard := board{
		found: [][]bool{
			{false, false, false, false, false},
			{true, true, true, true, true},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false}},
	}

	if !checkBingo(testBoard) {
		t.Error("Whoopsie daisy, func checkBingo is broken (horizontal bingo)")
	}
}

func TestCheckBingoNegative(t *testing.T) {
	testBoard := board{
		found: [][]bool{
			{false, false, false, false, false},
			{true, false, false, false, false},
			{true, false, false, false, false},
			{true, false, false, false, false},
			{true, false, false, false, false}},
	}

	if checkBingo(testBoard) {
		t.Error("Whoopsie daisy, func checkBingo is broken (negative)")
	}
}

func TestSumAllUnmarked(t *testing.T) {
	testBoard := board{
		found: [][]bool{
			{true, false, false, false, false},
			{true, true, false, false, true},
			{true, false, true, false, true},
			{true, false, false, true, false},
			{true, false, false, false, true}},
		values: [][]int{
			{0, 1, 1, 1, 1},
			{0, 0, 1, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 1, 0, 1},
			{0, 1, 1, 1, 0}},
	}

	expectedResult := 14
	actualResults := sumAllUnmarked(testBoard)

	if actualResults != expectedResult {
		t.Errorf("Invalid sum. Expected: %d. Actual: %d", expectedResult, actualResults)
	}
}
