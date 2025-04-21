package y2015_test

import (
	"testing"

	"github.com/ihribernik/advent-of-code/internal/solutions"
)

func TestDay05Part01(t *testing.T) {
	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult int
	}
	testCases := TestCase{
		{desc: "one long nice string", input: []string{"ugknbfddgicrmopn"}, expectedResult: 1},
		{desc: "one short nince string", input: []string{"aaa"}, expectedResult: 1},
		{desc: "one long naughty string", input: []string{"jchzalrnumimnmhp"}, expectedResult: 0},
		{desc: "second long naughty string", input: []string{"haegwjzuvuyypxyu"}, expectedResult: 0},
		{desc: "another long naughty string", input: []string{"dvszwmarrgswjxmb"}, expectedResult: 0},
	}

	solver, ok := solutions.GetSolver(2015, 05)
	if !ok {
		t.Errorf("failed to get solver")
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart1(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf(`solver.SolverPart1(%v) = %v, whants %v, error %v`, tC.input, result, tC.expectedResult, err)

			}
		})
	}
}

func TestDay05Part02(t *testing.T) {
	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult int
	}
	// For example:

	// qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).
	// xxyxx is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
	// uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a single letter between them.
	// ieodomkazucvgmuy is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.

	testCases := TestCase{
		{desc: "one long nice string", input: []string{"qjhvhtzxzqqjkmpb"}, expectedResult: 1},
		{desc: "one short nince string", input: []string{"xxyxx"}, expectedResult: 1},
		{desc: "one long naughty string", input: []string{"uurcxstgmygtbstg"}, expectedResult: 0},
		{desc: "second long naughty string", input: []string{"ieodomkazucvgmuy"}, expectedResult: 0},
	}

	solver, ok := solutions.GetSolver(2015, 05)
	if !ok {
		t.Errorf("failed to get solver")
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart2(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf(`solver.SolverPart1(%v) = %v, whants %v, error %v`, tC.input, result, tC.expectedResult, err)

			}
		})
	}
}
