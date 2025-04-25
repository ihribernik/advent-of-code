package y2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ihribernik/advent-of-code/internal/solutions"
)

type Day06 struct{}

func (d Day06) getMatrix() map[string]bool {
	matrix := map[string]bool{}
	for dx := range 1000 {
		for dy := range 1000 {
			key := fmt.Sprintf("%d,%d", dx, dy)
			matrix[key] = false
		}
	}
	return matrix
}

func (d Day06) getMatrixNumeric() map[string]int {
	matrix := map[string]int{}
	for dx := range 1000 {
		for dy := range 1000 {
			key := fmt.Sprintf("%d,%d", dx, dy)
			matrix[key] = 0
		}
	}
	return matrix
}

// SolvePart1 implements solutions.Solver.
func (d Day06) SolvePart1(input []string) (solutions.Solution, error) {
	result := 0
	matrix := d.getMatrix()
	for _, v := range input {
		containsGroups := regexp.MustCompile(`^(turn on|turn off|toggle) (\d+,\d+) through (\d+,\d+)$`)
		matches := containsGroups.FindStringSubmatch(v)
		action := matches[1]
		from := strings.Split(matches[2], ",")
		fromDx, err := strconv.Atoi(from[0])
		if err != nil {
			panic("not a valid from")
		}
		fromDy, err := strconv.Atoi(from[1])
		if err != nil {
			panic("not a valid from")
		}
		to := strings.Split(matches[3], ",")
		toDx, err := strconv.Atoi(to[0])
		if err != nil {
			panic("not a valid to")
		}
		toDy, err := strconv.Atoi(to[1])
		if err != nil {
			panic("not a valid to")
		}

		for x := fromDx; x <= toDx; x++ {
			for y := fromDy; y <= toDy; y++ {
				// fmt.Printf("x:%v, fromdx:%v, todx:%v, y:%d,fromDy:%d, toDy:%d\n", x, fromDx, toDx, y, fromDy, toDy)
				key := fmt.Sprintf("%v,%v", x, y)
				prevAction := matrix[key]
				posibleAction := false
				switch action {
				case "toggle":
					posibleAction = !prevAction
				case "turn off":
					posibleAction = false
				case "turn on":
					posibleAction = true
				default:
					panic("not a valid action")
				}
				matrix[key] = posibleAction
			}
		}
	}

	for _, status := range matrix {
		if status {
			result++
		}
	}
	return solutions.Solution{Result: result}, nil
}

// SolvePart2 implements solutions.Solver.
func (d Day06) SolvePart2(input []string) (solutions.Solution, error) {
	result := 0
	matrix := d.getMatrixNumeric()
	for _, v := range input {
		containsGroups := regexp.MustCompile(`^(turn on|turn off|toggle) (\d+,\d+) through (\d+,\d+)$`)
		matches := containsGroups.FindStringSubmatch(v)
		action := matches[1]
		from := strings.Split(matches[2], ",")
		fromDx, err := strconv.Atoi(from[0])
		if err != nil {
			panic("not a valid from")
		}
		fromDy, err := strconv.Atoi(from[1])
		if err != nil {
			panic("not a valid from")
		}
		to := strings.Split(matches[3], ",")
		toDx, err := strconv.Atoi(to[0])
		if err != nil {
			panic("not a valid to")
		}
		toDy, err := strconv.Atoi(to[1])
		if err != nil {
			panic("not a valid to")
		}

		for x := fromDx; x <= toDx; x++ {
			for y := fromDy; y <= toDy; y++ {
				// fmt.Printf("x:%v, fromdx:%v, todx:%v, y:%d,fromDy:%d, toDy:%d\n", x, fromDx, toDx, y, fromDy, toDy)
				key := fmt.Sprintf("%v,%v", x, y)
				prevBright := matrix[key]
				posibleBright := 0
				switch action {
				case "toggle":
					posibleBright = prevBright + 2
				case "turn off":
					if prevBright > 0 {
						posibleBright = prevBright - 1
					}
				case "turn on":
					posibleBright = prevBright + 1
				default:
					panic("not a valid action")
				}
				matrix[key] = posibleBright
			}
		}
	}

	for _, brightness := range matrix {
		result += brightness
	}

	return solutions.Solution{Result: result}, nil
}

func init() {
	solutions.Register(2015, 06, Day06{})
}
