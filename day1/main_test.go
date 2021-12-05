package main

import (
	"testing"
)

func TestCounting(t *testing.T) {
	testInput := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expectedResult := 7

	actualResult := CountIncreases(testInput)

	if expectedResult != actualResult {
		t.Log("Failed to count increases")
		t.Fatal()
	}
}

func TestSlidingWindow(t *testing.T) {
	testInput := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expectedResult := 5

	actualResult := CountSlidingWindowIncreases(testInput)

	if expectedResult != actualResult {
		t.Log("Failed to count sliding windows")
		t.Logf("Actual result: %d", actualResult)
		t.Fatal()
	}
}
