package y2015

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day02 struct{}

func (d Day02) SolvePart1(input []string) (int, error) {
	result := 0

	for i, val := range input {
		l, w, h, err := d.parseLine(val)
		if err != nil {
			return 0, newInputError(2, i+1, val, err)
		}

		areas := []int{l * w, w * h, h * l}
		areaTotal := 2 * (areas[0] + areas[1] + areas[2])
		extraPaper := slices.Min(areas)
		resultArea := areaTotal + extraPaper
		result += resultArea
	}
	return result, nil
}

func (d Day02) SolvePart2(input []string) (int, error) {
	result := 0

	for i, val := range input {
		l, w, h, err := d.parseLine(val)
		if err != nil {
			return 0, newInputError(2, i+1, val, err)
		}
		numbers := []int{l, w, h}
		slices.Sort(numbers)
		min1 := numbers[0]
		min2 := numbers[1]
		wrapFeet := min1 + min1 + min2 + min2
		bowFeet := l * w * h
		totalFeet := wrapFeet + bowFeet
		result += totalFeet
	}

	return result, nil
}

func (d Day02) parseLine(input string) (int, int, int, error) {
	inputVars := strings.Split(input, "x")
	if len(inputVars) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid box format %q: expected LxWxH", input)
	}

	l, err := strconv.Atoi(inputVars[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid length %q: %w", inputVars[0], err)
	}
	w, err := strconv.Atoi(inputVars[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid width %q: %w", inputVars[1], err)
	}
	h, err := strconv.Atoi(inputVars[2])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid height %q: %w", inputVars[2], err)
	}

	return l, w, h, nil
}
