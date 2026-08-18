package run

import (
	"slices"
)

type LoadInputFunc func(year, day int) ([]string, error)

type Result struct {
	Part1 int
	Part2 int
}

type runner struct {
	solvers   SolverResolver
	loadInput LoadInputFunc
}

// NewRunner constructs a runner with its required dependencies.
func NewRunner(solvers SolverResolver, loadInput LoadInputFunc) (Runner, error) {
	if solvers == nil {
		return nil, ConfigurationError{Dependency: "solver resolver", Err: ErrRunnerNotConfigured}
	}
	if loadInput == nil {
		return nil, ConfigurationError{Dependency: "input loader", Err: ErrRunnerNotConfigured}
	}

	return &runner{
		solvers:   solvers,
		loadInput: loadInput,
	}, nil
}

func (r *runner) Execute(year int, day int) (Result, error) {
	solver, ok := r.solvers.GetSolver(year, day)
	if !ok || solver == nil {
		return Result{}, &ErrSolverNotFound{Year: year, Day: day}
	}

	input, err := r.loadInput(year, day)
	if err != nil {
		return Result{}, &ErrGetInput{Year: year, Day: day, Err: err}
	}

	part1, err := solver.SolvePart1(slices.Clone(input))
	if err != nil {
		return Result{}, &ErrSolvePart{Part: 1, Err: err}
	}

	part2, err := solver.SolvePart2(slices.Clone(input))
	if err != nil {
		return Result{}, &ErrSolvePart{Part: 2, Err: err}
	}

	return Result{
		Part1: part1,
		Part2: part2,
	}, nil
}
