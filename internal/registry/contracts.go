package registry

type Solver interface {
	SolvePart1(input []string) (int, error)
	SolvePart2(input []string) (int, error)
}

type Registry interface {
	 Register(year int, day int, solver Solver) error
	 GetSolver(year int, day int) (Solver, bool)
}
