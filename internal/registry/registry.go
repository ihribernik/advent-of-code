package registry

import (
	"fmt"
)

type Key struct {
	Year int
	Day  int
}

type Registry struct {
	solvers map[Key]Solver
}

func NewRegistry() *Registry {
	return &Registry{
		solvers: make(map[Key]Solver),
	}
}

func (r *Registry) Register(year int, day int, solver Solver) error {
	if r == nil {
		return ErrNilRegistry
	}
	if solver == nil {
		return fmt.Errorf("nil solver for %d %d", year, day)
	}
	if r.solvers == nil {
		r.solvers = make(map[Key]Solver)
	}

	key := Key{year, day}
	if _, exists := r.solvers[key]; exists {
		return fmt.Errorf("solver already registered for %v", key)
	}

	r.solvers[key] = solver
	return nil
}

func (r *Registry) GetSolver(year int, day int) (Solver, bool) {
	if r == nil || r.solvers == nil {
		return nil, false
	}

	key := Key{year, day}
	solver, ok := r.solvers[key]
	return solver, ok
}
