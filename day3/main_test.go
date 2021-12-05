package main

import "testing"

func TestPowerConsumption(t *testing.T) {
	testInput := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	expectedResult := 198

	actualResult := powerConsumption(testInput)

	if actualResult != expectedResult {
		t.Error()
	}
}

func TestLifeSupportRating(t *testing.T) {
	testInput := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	expectedResult := 230

	actualResult := lifeSupportRating(testInput)

	if actualResult != expectedResult {
		t.Error()
	}
}

func TestOxygenRating(t *testing.T) {
	testInput := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	expectedResult := "10111"

	actualResult := findRating(testInput, 0, ">")

	if actualResult != expectedResult {
		t.Error()
	}
}

func TestCO2Rating(t *testing.T) {
	testInput := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	expectedResult := "01010"

	actualResult := findRating(testInput, 0, "<")

	if actualResult != expectedResult {
		t.Error()
	}
}
