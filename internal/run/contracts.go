package run

import "github.com/ihribernik/aoc-cli/internal/registry"

type Runner interface {
	Execute(year int, day int) (Result, error)
}

type SolverResolver interface {
	GetSolver(year int, day int) (registry.Solver, bool)
}
