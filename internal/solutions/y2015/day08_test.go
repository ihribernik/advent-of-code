package y2015_test

import (
	"testing"
)

func TestDay08Part1(t *testing.T) {
	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult int
	}

	testCases := TestCase{
		{
			desc: "empty string",
			input: []string{
				`""`,
				`"abc"`,
				`"aaa\"aaa"`,
				`"\x27"`,
			},
			expectedResult: 12,
		},
	}
	solver := mustSolver(t, 8)
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart1(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf("expected: %v, got: %v, errors: %v", tC.expectedResult, result, err)
			}
		})
	}
}

func TestDay08Part2(t *testing.T) {
	type TestCase = []struct {
		desc           string
		input          []string
		expectedResult int
	}

	testCases := TestCase{
		{
			desc: "empty string",
			input: []string{
				`""`,
				`"abc"`,
				`"aaa\"aaa"`,
				`"\x27"`,
			},
			expectedResult: 19,
		},
	}
	solver := mustSolver(t, 8)
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart2(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf("expected: %v, got: %v, errors: %v", tC.expectedResult, result, err)
			}
		})
	}
}
