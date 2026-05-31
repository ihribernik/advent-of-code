package y2015_test

import (
	"testing"
)

func TestDay04Part01(t *testing.T) {
	type TestCase = struct {
		desc           string
		input          []string
		expectedResult int
	}

	testCases := []TestCase{
		{
			desc: "answer is 609043...",
			input: []string{
				"abcdef",
			},
			expectedResult: 609043,
		},
		{
			desc: "answer is 1048970...",
			input: []string{
				"pqrstuv",
			},
			expectedResult: 1048970,
		},
	}

	solver := mustSolver(t, 0o4)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart1(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf(`solver.SolverPart1(%v) = %v, wants %v, error %v`, tC.input, result, tC.expectedResult, err)
			}
		})
	}
}

func TestDay04Part02(t *testing.T) {
	type TestCase = struct {
		desc           string
		input          []string
		expectedResult int
	}
	testCases := []TestCase{
		{
			desc:           "answer as six zeroes",
			input:          []string{"bgvyzdsv"},
			expectedResult: 1038736,
		},
	}

	solver := mustSolver(t, 0o4)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := solver.SolvePart2(tC.input)
			if result != tC.expectedResult || err != nil {
				t.Errorf(`solver.SolverPart2(%v) = %v, wants %v, error %v`, tC.input, result, tC.expectedResult, err)
			}
		})
	}
}
