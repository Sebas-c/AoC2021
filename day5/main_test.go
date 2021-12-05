package main

import "testing"

func TestPopulateMatrixWithDiagonals(t *testing.T) {
	testMatrix := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2"}

	expectResult := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 1, 2},
		{0, 1, 0, 0, 1, 0, 0, 1, 0, 2},
		{1, 1, 2, 0, 1, 0, 1, 0, 0, 2},
		{0, 1, 0, 1, 2, 1, 0, 0, 0, 1},
		{0, 0, 1, 0, 3, 0, 0, 0, 0, 1},
		{0, 0, 0, 2, 1, 2, 0, 0, 0, 1},
		{0, 0, 1, 0, 3, 0, 1, 0, 0, 0},
		{1, 2, 1, 2, 2, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0}}

	actualResult := populateMatrix(testMatrix, 10, true)

	for i := range expectResult {
		for j := range expectResult[i] {
			if expectResult[i][j] != actualResult[i][j] {
				t.Error()
			}
		}
	}
}

func TestPopulateMatrix(t *testing.T) {
	testMatrix := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2"}

	expectResult := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 2},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 2, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0}}

	actualResult := populateMatrix(testMatrix, 10, false)

	for i := range expectResult {
		for j := range expectResult[i] {
			if expectResult[i][j] != actualResult[i][j] {
				t.Error()
			}
		}
	}
}

func TestCountDoubleOverlap(t *testing.T) {
	testMatrix := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 2},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 2, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0}}

	expectResult := 5

	actualResult := countDoubleOverlap(testMatrix)

	if expectResult != actualResult {
		t.Error()
	}
}
