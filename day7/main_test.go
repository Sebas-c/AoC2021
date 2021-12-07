package main

import "testing"

func TestFindPosition(t *testing.T) {
	testInput := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	expectedResult := 37

	actualResult := findPosition(testInput)

	if expectedResult != actualResult {
		t.Errorf("Expected: %d, Actual: %d", expectedResult, actualResult)
	}
}

func TestFindExtremes(t *testing.T) {
	testInput := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	expectedResultHighest, expectedResultLowest := 16, 0

	actualResultHighest, actualResultLowest := findExtremes(testInput)

	if expectedResultHighest != actualResultHighest || expectedResultLowest != actualResultLowest {
		t.Errorf("Expected highest: %d, Actual: %d\nExpected lowest: %d, Actual: %d\n", expectedResultHighest, actualResultHighest, expectedResultLowest, actualResultLowest)
	}
}

func TestSumTo(t *testing.T) {
	testInput := 11

	expectedResult := 66

	actualResult := sumTo(testInput)

	if expectedResult != actualResult {
		t.Errorf("Expected: %d, Actual: %d", expectedResult, actualResult)
	}
}

func TestFindPositionAgain(t *testing.T) {
	testInput := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	expectedResult := 168

	actualResult := findPositionAgain(testInput)

	if expectedResult != actualResult {
		t.Errorf("Expected: %d, Actual: %d", expectedResult, actualResult)
	}
}
