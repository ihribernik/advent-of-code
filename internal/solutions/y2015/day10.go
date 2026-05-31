package y2015

import (
	"strconv"
	"strings"
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
	entries := strings.TrimSpace(input[0])

	for range 40 {
		entries = lookAndSay(entries)
	}

	return len(entries), nil
}

func (d Day10) SolvePart2(input []string) (int, error) {
	panic("not implemented") // TODO: Implement
}
