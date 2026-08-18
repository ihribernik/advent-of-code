package y2015

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ihribernik/aoc-cli/internal/puzzle"
)

type Day10 struct{}

type Run struct {
	Digit rune
	Count int
}

func lookAndSay(seq string) string {
	runs := ParseRun(seq)
	return RuntToString(runs)
}

func RuntToString(runs []Run) string {
	var b strings.Builder

	for _, run := range runs {
		b.WriteString(strconv.Itoa(run.Count))
		b.WriteRune(run.Digit)
	}

	return b.String()
}

func ParseRun(seq string) []Run {
	var runs []Run

	return runs
}

func (d Day10) SolvePart1(input []string) (int, error) {
	if len(input) == 0 || strings.TrimSpace(input[0]) == "" {
		return 0, newInputError(10, 0, "", errors.New("missing starting sequence"))
	}

	return 0, fmt.Errorf("2015 day 10 part 1: %w", puzzle.ErrNotImplemented)
}

func (d Day10) SolvePart2(input []string) (int, error) {
	if len(input) == 0 || strings.TrimSpace(input[0]) == "" {
		return 0, newInputError(10, 0, "", errors.New("missing starting sequence"))
	}

	return 0, fmt.Errorf("2015 day 10 part 2: %w", puzzle.ErrNotImplemented)
}
