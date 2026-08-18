package y2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day06 struct{}

var day06InstructionPattern = regexp.MustCompile(`^(turn on|turn off|toggle) (\d+,\d+) through (\d+,\d+)$`)

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

func parseCoordPair(raw string) (int, int, error) {
	parts := strings.Split(raw, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid coordinate pair %q", raw)
	}

	dx, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid x coordinate %q: %w", parts[0], err)
	}
	dy, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid y coordinate %q: %w", parts[1], err)
	}
	if dx < 0 || dx > 999 || dy < 0 || dy > 999 {
		return 0, 0, fmt.Errorf("coordinate %q is outside the 0-999 grid", raw)
	}

	return dx, dy, nil
}

func parseInstruction(line string) (action string, fromDx int, fromDy int, toDx int, toDy int, err error) {
	matches := day06InstructionPattern.FindStringSubmatch(line)
	if matches == nil {
		return "", 0, 0, 0, 0, fmt.Errorf("invalid instruction %q", line)
	}

	fromDx, fromDy, err = parseCoordPair(matches[2])
	if err != nil {
		return "", 0, 0, 0, 0, err
	}
	toDx, toDy, err = parseCoordPair(matches[3])
	if err != nil {
		return "", 0, 0, 0, 0, err
	}
	if fromDx > toDx || fromDy > toDy {
		return "", 0, 0, 0, 0, fmt.Errorf("coordinate range must start before it ends")
	}

	return matches[1], fromDx, fromDy, toDx, toDy, nil
}

// SolvePart1 implements solutions.Solver.
func (d Day06) SolvePart1(input []string) (int, error) {
	result := 0
	matrix := d.getMatrix()

	for i, v := range input {
		action, fromDx, fromDy, toDx, toDy, err := parseInstruction(v)
		if err != nil {
			return 0, newInputError(6, i+1, v, err)
		}

		for x := fromDx; x <= toDx; x++ {
			for y := fromDy; y <= toDy; y++ {
				key := fmt.Sprintf("%v,%v", x, y)
				prevAction := matrix[key]
				nextAction := false
				switch action {
				case "toggle":
					nextAction = !prevAction
				case "turn off":
					nextAction = false
				case "turn on":
					nextAction = true
				default:
					return 0, fmt.Errorf("invalid action %q", action)
				}
				matrix[key] = nextAction
			}
		}
	}

	for _, status := range matrix {
		if status {
			result++
		}
	}
	return result, nil
}

// SolvePart2 implements solutions.Solver.
func (d Day06) SolvePart2(input []string) (int, error) {
	result := 0
	matrix := d.getMatrixNumeric()

	for i, v := range input {
		action, fromDx, fromDy, toDx, toDy, err := parseInstruction(v)
		if err != nil {
			return 0, newInputError(6, i+1, v, err)
		}

		for x := fromDx; x <= toDx; x++ {
			for y := fromDy; y <= toDy; y++ {
				key := fmt.Sprintf("%v,%v", x, y)
				prevBright := matrix[key]
				nextBright := 0
				switch action {
				case "toggle":
					nextBright = prevBright + 2
				case "turn off":
					if prevBright > 0 {
						nextBright = prevBright - 1
					}
				case "turn on":
					nextBright = prevBright + 1
				default:
					return 0, fmt.Errorf("invalid action %q", action)
				}
				matrix[key] = nextBright
			}
		}
	}

	for _, brightness := range matrix {
		result += brightness
	}

	return result, nil
}
