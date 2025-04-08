package solutions

import (
	"fmt"
)

type Solver interface {
	SolvePart1(input []string) (int, error)
	SolvePart2(input []string) (int, error)
}

var registry = make(map[string]Solver)

func Register(year int, day int, solver Solver) {
	key := fmt.Sprintf("%d-%02d", year, day)
	registry[key] = solver
}

func GetSolver(year int, day int) (Solver, bool) {
	key := fmt.Sprintf("%d-%02d", year, day)
	solver, ok := registry[key]
	return solver, ok
}
