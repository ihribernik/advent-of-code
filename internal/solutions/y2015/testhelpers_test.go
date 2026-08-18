package y2015_test

import (
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
	"github.com/ihribernik/aoc-cli/internal/solutions/y2015"
)

const testYear = 2015

func mustSolver(t *testing.T, day int) registry.Solver {
	t.Helper()

	reg := registry.NewRegistry()
	if err := y2015.Register(reg); err != nil {
		t.Fatalf("failed to load registry for the year %d, err: %v", testYear, err)
	}

	solver, ok := reg.GetSolver(testYear, day)
	if !ok {
		t.Fatalf("failed to get solver for year %d day %d", testYear, day)
	}

	return solver
}
