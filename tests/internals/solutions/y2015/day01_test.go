package y2015_test

import (
	"testing"

	"github.com/ihribernik/advent-of-code/internal/solutions"
	_ "github.com/ihribernik/advent-of-code/internal/solutions/y2015"
)

func TestDay01_SolvePart1(t *testing.T) {

	type TestCase struct {
		input          []string
		expectedResult int
	}

	testsCases := []TestCase{
		{[]string{"("}, 1},
		{[]string{"(", ")"}, 0},
		{[]string{"(", "(", ")"}, 1},
		{[]string{"(", ")", "(", ")"}, 0},
		{[]string{"(", "(", ")", "(", ")"}, 1},
	}

	solver, ok := solutions.GetSolver(2015, 01)
	if !ok {
		t.Errorf("failed to get solver")
	}

	for i, testCase := range testsCases {
		testCase := testCase
		t.Run("", func(t *testing.T) {
			result, err := solver.SolvePart1(testCase.input)

			if err != nil {

				if result != testCase.expectedResult {
					t.Fatalf("test %v: expected: %v, got: %v", i+1, testCase.expectedResult, result)
				}
			}
		})

	}

}

func TestDay01_SolvePart2(t *testing.T) {

	type TestCase struct {
		input    []string
		expected int
	}

	testsCases := []TestCase{
		{[]string{"("}, 1},
		{[]string{"(", ")"}, 0},
		{[]string{"(", "(", ")"}, 1},
		{[]string{"(", ")", "(", ")"}, 0},
		{[]string{"(", "(", ")", "(", ")"}, 1},
	}

	solver, ok := solutions.GetSolver(2015, 01)
	if !ok {
		t.Errorf("failed to get solver")
	}

	for i, testCase := range testsCases {
		testCase := testCase
		t.Run("", func(t *testing.T) {
			result, err := solver.SolvePart2(testCase.input)

			if err != nil {

				if result != testCase.expected {
					t.Fatalf("test %v: expected: %v, got: %v", i+1, testCase.expected, result)
				}
			}
		})

	}

}
