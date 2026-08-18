package y2015_test

import (
	"errors"
	"testing"

	"github.com/ihribernik/aoc-cli/internal/puzzle"
	"github.com/ihribernik/aoc-cli/internal/solutions/y2015"
)

func TestDay10ReturnsNotImplemented(t *testing.T) {
	solver := y2015.Day10{}

	tests := []struct {
		name  string
		solve func([]string) (int, error)
	}{
		{name: "part 1", solve: solver.SolvePart1},
		{name: "part 2", solve: solver.SolvePart2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.solve([]string{"1"})
			if !errors.Is(err, puzzle.ErrNotImplemented) {
				t.Fatalf("expected ErrNotImplemented, got %v", err)
			}
		})
	}
}

func TestDay10RejectsMissingInput(t *testing.T) {
	solver := y2015.Day10{}

	for _, solve := range []func([]string) (int, error){solver.SolvePart1, solver.SolvePart2} {
		_, err := solve(nil)
		if !errors.Is(err, puzzle.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	}
}
