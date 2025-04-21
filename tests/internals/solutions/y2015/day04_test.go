package y2015_test

import (
	"testing"

	"github.com/ihribernik/advent-of-code/internal/solutions"
)

// --- Day 4: The Ideal Stocking Stuffer ---
// Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts for all the economically forward-thinking little girls and boys.

// To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least five zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal.
// To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.

// For example:

// If your secret key is abcdef, the answer is 609043, because the MD5 hash of abcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest such number to do so.
// If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....

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
			desc: "anwser is 1048970...",
			input: []string{
				"pqrstuv",
			},
			expectedResult: 1048970,
		},
	}

	solver, ok := solutions.GetSolver(2015, 04)

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
