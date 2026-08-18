package y2015

import "github.com/ihribernik/aoc-cli/internal/puzzle"

func newInputError(day int, line int, input string, err error) error {
	return &puzzle.InputError{Year: 2015, Day: day, Line: line, Input: input, Err: err}
}
