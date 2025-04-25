package y2015_test

import (
	"testing"

	"github.com/ihribernik/advent-of-code/internal/solutions"
	_ "github.com/ihribernik/advent-of-code/internal/solutions/y2015"
)

func TestDay01_SolvePart1(t *testing.T) {

	type TestCase struct {
		input          []string
		expectedResult solutions.Solution
	}

	testsCases := []TestCase{
		{[]string{"("}, solutions.Solution{Result: 1}},
		{[]string{"(", ")"}, solutions.Solution{Result: 0}},
		{[]string{"(", "(", ")"}, solutions.Solution{Result: 1}},
		{[]string{"(", ")", "(", ")"}, solutions.Solution{Result: 0}},
		{[]string{"(", "(", ")", "(", ")"}, solutions.Solution{Result: 1}},
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

				if result.Result != testCase.expectedResult.Result {
					t.Fatalf("test %v: expected: %v, got: %v", i+1, testCase.expectedResult, result)
				}
			}
		})

	}

}

func TestDay01_SolvePart2(t *testing.T) {

	type TestCase struct {
		input    []string
		expected solutions.Solution
	}

	testsCases := []TestCase{
		{[]string{"("}, solutions.Solution{Result: 1}},
		{[]string{"(", ")"}, solutions.Solution{Result: 0}},
		{[]string{"(", "(", ")"}, solutions.Solution{Result: 1}},
		{[]string{"(", ")", "(", ")"}, solutions.Solution{Result: 0}},
		{[]string{"(", "(", ")", "(", ")"}, solutions.Solution{Result: 1}},
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

				if result.Result != testCase.expected.Result {
					t.Fatalf("test %v: expected: %v, got: %v", i+1, testCase.expected, result)
				}
			}
		})

	}

}
