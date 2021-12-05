package main

import "testing"

func TestPosition(t *testing.T) {
	testInput := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	expectedResult := 150

	actualResult := calculatePosition(testInput)

	if actualResult != expectedResult {
		t.Logf("Actual result: %d", actualResult)
		t.Error()
	}
}

func TestPositionWithAim(t *testing.T) {
	testInput := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	expectedResult := 900

	actualResult := calculatePositionWithAim(testInput)

	if actualResult != expectedResult {
		t.Logf("Actual result: %d", actualResult)
		t.Error()
	}
}
