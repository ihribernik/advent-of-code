package y2015

import "github.com/ihribernik/advent-of-code/internal/solutions"

type Day01 struct{}

func (d Day01) SolvePart1(input []string) (solutions.Solution, error) {
	result := 0
	for _, v := range input {
		switch v {
		case "(":
			result += 1
		case ")":
			result -= 1

		default:
			continue
		}
	}

	return solutions.Solution{Result: result}, nil
}

func (d Day01) SolvePart2(input []string) (solutions.Solution, error) {
	positionChar := 0
	floor := 0

	for index, v := range input {
		if floor == -1 {
			positionChar = index
			break
		}

		switch v {
		case "(":
			floor += 1
		case ")":
			floor -= 1
		default:
			continue
		}
	}

	return solutions.Solution{Result: positionChar}, nil
}

func init() {
	solutions.Register(2015, 01, Day01{})
}
