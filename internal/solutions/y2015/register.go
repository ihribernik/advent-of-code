package y2015

import (
	"fmt"

	"github.com/ihribernik/aoc-cli/internal/registry"
)

type entry struct {
	day    int
	solver registry.Solver
}

func Register(r *registry.Registry) error {
	entries := []entry{
		{1, Day01{}},
		{2, Day02{}},
		{3, Day03{}},
		{4, Day04{}},
		{5, Day05{}},
		{6, Day06{}},
		{7, Day07{}},
		{8, Day08{}},
		{9, Day09{}},
	}

	for _, e := range entries {
		if err := r.Register(2015, e.day, e.solver); err != nil {
			return fmt.Errorf("2015 day %02d: %w", e.day, err)
		}
	}
	return nil
}
