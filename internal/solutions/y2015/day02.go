package y2015

import (
	"slices"
	"strconv"
	"strings"

	"github.com/ihribernik/advent-of-code/internal/solutions"
)

type Day02 struct{}

func (d Day02) SolvePart1(input []string) (solutions.Solution, error) {
	result := 0

	for _, val := range input {
		l, w, h, err := d.parseLine(val)
		if err != nil {
			return solutions.Solution{Result: 0}, err
		}

		areas := []int{l * w, w * h, h * l}
		area_total := 2 * (areas[0] + areas[1] + areas[2])
		extra_paper := slices.Min(areas)
		result_area := area_total + extra_paper
		result += result_area
	}
	return solutions.Solution{Result: result}, nil
}

func (d Day02) SolvePart2(input []string) (solutions.Solution, error) {
	result := 0

	for _, val := range input {
		l, w, h, err := d.parseLine(val)
		if err != nil {
			return solutions.Solution{Result: 0}, err
		}
		numbers := []int{l, w, h}
		slices.Sort(numbers)
		min_1 := numbers[0]
		min_2 := numbers[1]
		wrap_feet := min_1 + min_1 + min_2 + min_2
		bow_feet := l * w * h
		total_feet := wrap_feet + bow_feet
		result += total_feet
	}

	return solutions.Solution{Result: result}, nil

}

func (d Day02) parseLine(input string) (int, int, int, error) {
	inputVars := strings.Split(input, "x")
	l, err := strconv.Atoi(inputVars[0])

	if err != nil {
		return 0, 0, 0, err
	}
	w, err := strconv.Atoi(inputVars[1])
	if err != nil {
		return 0, 0, 0, err
	}
	h, err := strconv.Atoi(inputVars[2])
	if err != nil {
		return 0, 0, 0, err
	}

	return l, w, h, nil
}

func init() {
	solutions.Register(2015, 02, Day02{})
}
