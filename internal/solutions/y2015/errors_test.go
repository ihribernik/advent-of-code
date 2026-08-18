package y2015

import (
	"errors"
	"strconv"
	"testing"

	"github.com/ihribernik/aoc-cli/internal/puzzle"
)

func TestSolverInvalidInputErrors(t *testing.T) {
	tests := []struct {
		name  string
		solve func([]string) (int, error)
		input []string
		day   int
		line  int
	}{
		{name: "day 2 number", solve: Day02{}.SolvePart1, input: []string{"2x3x4", "2x3xnope"}, day: 2, line: 2},
		{name: "day 4 missing key", solve: Day04{}.SolvePart1, input: nil, day: 4, line: 0},
		{name: "day 6 malformed instruction", solve: Day06{}.SolvePart1, input: []string{"turnon 0,0 through 1,1"}, day: 6, line: 1},
		{name: "day 6 coordinate outside grid", solve: Day06{}.SolvePart1, input: []string{"turn on 0,0 through 1000,1"}, day: 6, line: 1},
		{name: "day 6 reversed range", solve: Day06{}.SolvePart1, input: []string{"turn on 2,2 through 1,1"}, day: 6, line: 1},
		{name: "day 7 malformed operation", solve: Day07{}.SolvePart1, input: []string{"bad operation"}, day: 7, line: 1},
		{name: "day 7 missing wire", solve: Day07{}.SolvePart1, input: []string{"x -> a"}, day: 7, line: 0},
		{name: "day 8 invalid quoted string", solve: Day08{}.SolvePart1, input: []string{"not quoted"}, day: 8, line: 1},
		{name: "day 9 malformed route", solve: Day09{}.SolvePart1, input: []string{"London Dublin 464"}, day: 9, line: 1},
		{name: "day 9 invalid distance", solve: Day09{}.SolvePart1, input: []string{"London to Dublin = far"}, day: 9, line: 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.solve(tc.input)
			if !errors.Is(err, puzzle.ErrInvalidInput) {
				t.Fatalf("expected ErrInvalidInput, got %v", err)
			}

			var inputErr *puzzle.InputError
			if !errors.As(err, &inputErr) {
				t.Fatalf("expected InputError, got %T", err)
			}
			if inputErr.Year != 2015 || inputErr.Day != tc.day || inputErr.Line != tc.line {
				t.Fatalf("unexpected location: day %d line %d", inputErr.Day, inputErr.Line)
			}
		})
	}
}

func TestInputErrorPreservesNumberCause(t *testing.T) {
	_, err := Day02{}.SolvePart1([]string{"2x3xnope"})

	var numberErr *strconv.NumError
	if !errors.As(err, &numberErr) {
		t.Fatalf("expected strconv.NumError, got %v", err)
	}
}

func TestFindHashWithPrefixReturnsAnswerNotFound(t *testing.T) {
	_, err := findHashWithPrefix("secret", "impossible-prefix", 1)
	if !errors.Is(err, puzzle.ErrAnswerNotFound) {
		t.Fatalf("expected ErrAnswerNotFound, got %v", err)
	}
}
