package y2015_test

import (
	"testing"

	"github.com/ihribernik/advent-of-code/internal/solutions"
)

func TestDay05Part01(t *testing.T) {
	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult solutions.Solution
	}
	testCases := TestCase{
		{desc: "one long nice string", input: []string{"ugknbfddgicrmopn"}, expectedResult: solutions.Solution{Result: 1}},
		{desc: "one short nince string", input: []string{"aaa"}, expectedResult: solutions.Solution{Result: 1}},
		{desc: "one long naughty string", input: []string{"jchzalrnumimnmhp"}, expectedResult: solutions.Solution{Result: 0}},
		{desc: "second long naughty string", input: []string{"haegwjzuvuyypxyu"}, expectedResult: solutions.Solution{Result: 0}},
		{desc: "another long naughty string", input: []string{"dvszwmarrgswjxmb"}, expectedResult: solutions.Solution{Result: 0}},
	}

	solver, ok := solutions.GetSolver(2015, 05)
	if !ok {
		t.Errorf("failed to get solver")
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart1(tC.input)
			if result.Result != tC.expectedResult.Result || err != nil {
				t.Errorf(`solver.SolverPart1(%v) = %v, whants %v, error %v`, tC.input, result, tC.expectedResult, err)

			}
		})
	}
}

func TestDay05Part02(t *testing.T) {
	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult solutions.Solution
	}

	testCases := TestCase{
		{desc: "one long nice string", input: []string{"qjhvhtzxzqqjkmpb"}, expectedResult: solutions.Solution{Result: 1}},
		{desc: "one short nince string", input: []string{"xxyxx"}, expectedResult: solutions.Solution{Result: 1}},
		{desc: "one long naughty string", input: []string{"uurcxstgmygtbstg"}, expectedResult: solutions.Solution{Result: 0}},
		{desc: "second long naughty string", input: []string{"ieodomkazucvgmuy"}, expectedResult: solutions.Solution{Result: 0}},
	}

	solver, ok := solutions.GetSolver(2015, 05)
	if !ok {
		t.Errorf("failed to get solver")
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart2(tC.input)
			if result.Result != tC.expectedResult.Result || err != nil {
				t.Errorf(`solver.SolverPart1(%v) = %v, whants %v, error %v`, tC.input, result, tC.expectedResult, err)

			}
		})
	}
}
