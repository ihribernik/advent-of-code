package y2015_test

import (
	"testing"

	"github.com/ihribernik/advent-of-code/internal/solutions"
	_ "github.com/ihribernik/advent-of-code/internal/solutions/y2015"
)

func TestDay03Part01(t *testing.T) {
	type TestCase = struct {
		desc           string
		input          []string
		expectedResult solutions.Solution
	}

	testCases := []TestCase{
		{
			desc: "delivers presents to 2 houses: one at the starting location, and one to the east",
			input: []string{
				">",
			},
			expectedResult: solutions.Solution{Result: 2},
		},
		{
			desc: "delivers presents to 4 houses in a square, including twice to the house at his starting/ending location",
			input: []string{
				"^>v<",
			},
			expectedResult: solutions.Solution{Result: 4},
		},
		{
			desc: "delivers a bunch of presents to some very lucky children at only 2 houses",
			input: []string{
				"^v^v^v^v^v",
			},
			expectedResult: solutions.Solution{Result: 2},
		},
	}

	solver, ok := solutions.GetSolver(2015, 03)
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

func TestDay03Part02(t *testing.T) {
	type TestCase = struct {
		desc           string
		input          []string
		expectedResult solutions.Solution
	}

	// ^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
	// ^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
	// ^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.
	testCases := []TestCase{
		{
			desc: "delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.",
			input: []string{
				"^v",
			},
			expectedResult: solutions.Solution{Result: 3},
		},
		{
			desc: "now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.",
			input: []string{
				"^>v<",
			},
			expectedResult: solutions.Solution{Result: 3},
		},
		{
			desc: "delivers a bunch of presents to some very lucky children at only 2 houses",
			input: []string{
				"^v^v^v^v^v",
			},
			expectedResult: solutions.Solution{Result: 11},
		},
	}

	solver, ok := solutions.GetSolver(2015, 03)
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
