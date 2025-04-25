package y2015_test

import (
	"testing"

	"github.com/ihribernik/advent-of-code/internal/solutions"
)

func TestDay06Part1(t *testing.T) {
	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult solutions.Solution
	}

	testCases := TestCase{
		{desc: "may turn on only 1000000", input: []string{"turn on 0,0 through 999,999"}, expectedResult: solutions.Solution{Result: 1000000}},
		{desc: "may turn on only 1000", input: []string{"toggle 0,0 through 999,0"}, expectedResult: solutions.Solution{Result: 1000}},
		{desc: "may turn on only 4", input: []string{"turn on 499,499 through 500,500"}, expectedResult: solutions.Solution{Result: 4}},
	}
	solver, ok := solutions.GetSolver(2015, 06)

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

func TestDay06Part2(t *testing.T) {
	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult solutions.Solution
	}

	testCases := TestCase{
		{desc: "may total brightness eq 1", input: []string{"turn on 0,0 through 0,0"}, expectedResult: solutions.Solution{Result: 1}},
		{desc: "may total brightness eq 2000000.", input: []string{"toggle 0,0 through 999,999"}, expectedResult: solutions.Solution{Result: 2000000}},
	}
	solver, ok := solutions.GetSolver(2015, 06)

	if !ok {
		t.Errorf("failed to get solver")
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart2(tC.input)
			if result.Result != tC.expectedResult.Result || err != nil {
				t.Errorf(`solver.SolverPart2(%v) = %v, whants %v, error %v`, tC.input, result, tC.expectedResult, err)
			}
		})
	}
}
