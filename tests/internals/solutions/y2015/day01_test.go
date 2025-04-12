package y2015

import "testing"

func TestDay01_SolvePart1(t *testing.T) {

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

	day01 := Day01{}

	for i, testCase := range testsCases {
		testCase := testCase
		t.Run("", func(t *testing.T) {
			result, err := day01.SolvePart1(testCase.input)

			if err != nil {

				if result != testCase.expected {
					t.Fatalf("test %v: expected: %v, got: %v", i+1, testCase.expected, result)
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

	day01 := Day01{}

	for i, testCase := range testsCases {
		testCase := testCase
		t.Run("", func(t *testing.T) {
			result, err := day01.SolvePart2(testCase.input)

			if err != nil {

				if result != testCase.expected {
					t.Fatalf("test %v: expected: %v, got: %v", i+1, testCase.expected, result)
				}
			}
		})

	}

}
