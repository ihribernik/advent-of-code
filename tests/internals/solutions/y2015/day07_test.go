package y2015_test

import (
	"testing"

	"github.com/ihribernik/advent-of-code/internal/solutions"
)

func TestDay07Part1(t *testing.T) {

	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult solutions.Solution
	}

	testCases := TestCase{
		{
			desc: "test circuit",
			input: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			expectedResult: solutions.Solution{
				MapResult: map[string]any{
					"d": 72,
					"e": 507,
					"f": 492,
					"g": 114,
					"h": 65412,
					"i": 65079,
					"x": 123,
					"y": 456,
				},
			},
		},
	}
	solver, ok := solutions.GetSolver(2015, 07)

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
