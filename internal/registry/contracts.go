package registry

type Solver interface {
	SolvePart1(input []string) (int, error)
	SolvePart2(input []string) (int, error)
}
