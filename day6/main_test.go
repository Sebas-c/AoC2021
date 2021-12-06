package main

import "testing"

func TestFishCount(t *testing.T) {
	testArray := []int{3, 4, 3, 1, 2}
	days := 80

	expectResult := 5934

	actualResult := fishCount(testArray, days)

	if expectResult != actualResult {
		t.Logf("Expected: %d, Actual: %d", expectResult, actualResult)
		t.Error()
	}
}

func TestFishCountTheSmartWay(t *testing.T) {
	testArray := []int{3, 4, 3, 1, 2}
	days := 256

	expectResult := 26984457539

	actualResult := fishCountTheSmartWay(testArray, days)

	if expectResult != actualResult {
		t.Logf("Expected: %d, Actual: %d", expectResult, actualResult)
		t.Error()
	}
}
