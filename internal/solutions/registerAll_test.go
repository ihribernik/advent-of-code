package solutions_test

import (
	"errors"
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
	"github.com/ihribernik/aoc-cli/internal/solutions"
)

func TestRegisterAll(t *testing.T) {
	t.Run("registers all solvers", func(t *testing.T) {
		r := registry.NewRegistry()
		if err := solutions.RegisterAll(r); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if solver, ok := r.GetSolver(2015, 1); !ok || solver == nil {
			t.Fatal("expected solver for 2015 day 1")
		}
		if solver, ok := r.GetSolver(2015, 10); !ok || solver == nil {
			t.Fatal("expected solver for 2015 day 10")
		}
	})

	t.Run("rejects nil registry", func(t *testing.T) {
		err := solutions.RegisterAll(nil)
		if !errors.Is(err, registry.ErrNilRegistry) {
			t.Fatalf("expected ErrNilRegistry, got %v", err)
		}
	})

	t.Run("preserves duplicate registration details", func(t *testing.T) {
		r := registry.NewRegistry()
		if err := solutions.RegisterAll(r); err != nil {
			t.Fatalf("unexpected setup error: %v", err)
		}

		err := solutions.RegisterAll(r)
		if !errors.Is(err, registry.ErrSolverAlreadyRegistered) {
			t.Fatalf("expected duplicate registration sentinel, got %v", err)
		}
		var registrationErr *registry.RegistrationError
		if !errors.As(err, &registrationErr) {
			t.Fatalf("expected RegistrationError, got %T", err)
		}
		if registrationErr.Year != 2015 || registrationErr.Day != 1 {
			t.Fatalf("unexpected registration metadata: %+v", registrationErr)
		}
	})
}
