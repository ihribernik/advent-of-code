package y2015

import (
	"strconv"
)

type Day08 struct{}

// SolvePart1 implements solutions.Solver.
func (d Day08) SolvePart1(input []string) (int, error) {
	inMemoryCounter := 0
	literalsCounter := 0

	for i, line := range input {
		if len(line) == 0 {
			continue
		}

		literalsCounter += len(line)
		evaluated, err := strconv.Unquote(line)
		if err != nil {
			return 0, newInputError(8, i+1, line, err)
		}

		inMemoryCounter += len(evaluated)

	}

	total := literalsCounter - inMemoryCounter
	return total, nil
}

// SolvePart2 implements solutions.Solver.
func (d Day08) SolvePart2(input []string) (int, error) {
	quotedLiteralsCounter := 0
	literalsCounter := 0

	for _, line := range input {
		if len(line) == 0 {
			continue
		}

		literalsCounter += len(line)
		evaluated := strconv.Quote(line)
		quotedLiteralsCounter += len(evaluated)

	}

	total := quotedLiteralsCounter - literalsCounter
	return total, nil
}
