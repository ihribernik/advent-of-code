package solutions

import (
	"fmt"

	"github.com/ihribernik/aoc-cli/internal/registry"
	"github.com/ihribernik/aoc-cli/internal/solutions/y2015"
)

// RegisterAll registers every available solver.
func RegisterAll(r registry.Registry) error {
	if r == nil {
		return registry.ErrNilRegistry
	}

	if err := y2015.Register(r); err != nil {
		return fmt.Errorf("register year 2015: %w", err)
	}

	return nil
}
