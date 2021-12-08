package main

import (
	"strings"
	"testing"
)

func TestFindUniqueDigits(t *testing.T) {
	testInput := []string{"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"}

	expectedResult := 26

	actualResult := findUniqueDigits(testInput)

	if expectedResult != actualResult {
		t.Errorf("Expected result: %d, Actual result: %d", expectedResult, actualResult)
	}
}

func TestDecypherPatterns(t *testing.T) {
	testInput := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}

	expectedResult := []string{"cagedb", "ab", "gcdfa", "fbcad", "eafb", "cdfbe", "cdfgeb", "dab", "acedgfb", "cefabd"}

	actualResult := decypherPatterns(testInput)

	for i := range actualResult {
		if actualResult[i] != expectedResult[i] {
			t.Errorf("Expected result: %s, Actual result: %s", expectedResult[i], actualResult[i])
		}
	}
}

func TestSortAlphabetically(t *testing.T) {
	testInput := "acedgfb"

	expectedResult := "abcdefg"

	actualResult := sortAlphabetically(testInput)

	if strings.Compare(actualResult, expectedResult) != 0 {
		t.Errorf("Expected result: %s, Actual result: %s", expectedResult, actualResult)
	}
}

func TestFindTotalOutput(t *testing.T) {
	testInput := []string{"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"}

	expectedResult := 61229

	actualResult := findTotalOutput(testInput)

	if expectedResult != actualResult {
		t.Errorf("Expected result: %d, Actual result: %d", expectedResult, actualResult)
	}
}
