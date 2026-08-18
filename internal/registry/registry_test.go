package registry_test

import (
	"errors"
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
)

type testSolver struct{}

func (s testSolver) SolvePart1(input []string) (int, error) { return len(input), nil }
func (s testSolver) SolvePart2(input []string) (int, error) { return len(input), nil }

func TestNewRegistry(t *testing.T) {
	r := registry.NewRegistry()
	if r == nil {
		t.Fatalf("expected non-nil registry")
	}

	if _, ok := r.GetSolver(2015, 1); ok {
		t.Fatalf("expected empty registry")
	}
}

func TestRegistryRegisterAndGetSolver(t *testing.T) {
	r := registry.NewRegistry()
	s := testSolver{}

	if err := r.Register(2015, 1, s); err != nil {
		t.Fatalf("unexpected register error: %v", err)
	}

	got, ok := r.GetSolver(2015, 1)
	if !ok {
		t.Fatalf("expected solver to exist")
	}
	if got == nil {
		t.Fatalf("expected non-nil solver")
	}
}

func TestRegistryRegisterNilSolver(t *testing.T) {
	r := registry.NewRegistry()

	err := r.Register(2015, 1, nil)
	if !errors.Is(err, registry.ErrNilSolver) {
		t.Fatalf("expected ErrNilSolver, got %v", err)
	}
	assertRegistrationError(t, err, 2015, 1)
}

func TestRegistryRegisterDuplicate(t *testing.T) {
	r := registry.NewRegistry()
	s := testSolver{}

	if err := r.Register(2015, 1, s); err != nil {
		t.Fatalf("unexpected register error: %v", err)
	}
	if err := r.Register(2015, 1, s); !errors.Is(err, registry.ErrSolverAlreadyRegistered) {
		t.Fatalf("expected ErrSolverAlreadyRegistered, got %v", err)
	} else {
		assertRegistrationError(t, err, 2015, 1)
	}
}

func assertRegistrationError(t *testing.T, err error, year, day int) {
	t.Helper()
	var registrationErr *registry.RegistrationError
	if !errors.As(err, &registrationErr) {
		t.Fatalf("expected RegistrationError, got %T", err)
	}
	if registrationErr.Year != year || registrationErr.Day != day {
		t.Fatalf("unexpected registration metadata: %+v", registrationErr)
	}
}

func TestRegistryRegisterInitializesNilMap(t *testing.T) {
	r := registry.NewRegistry()
	s := testSolver{}

	if err := r.Register(2015, 1, s); err != nil {
		t.Fatalf("unexpected register error: %v", err)
	}

	if got, ok := r.GetSolver(2015, 1); !ok || got == nil {
		t.Fatalf("expected solver after lazy map initialization")
	}
}

func TestRegistryGetSolverOnNilMap(t *testing.T) {
	r := registry.NewRegistry()

	if got, ok := r.GetSolver(2015, 1); ok || got != nil {
		t.Fatalf("expected nil,false for nil map, got %v,%v", got, ok)
	}
}
