package y2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ihribernik/advent-of-code/internal/solutions"
)

// --- Day 6: Probably a Fire Hazard ---

// Because your neighbors keep defeating you in the holiday house decorating contest year after year, you've decided to deploy one million lights in a 1000x1000 grid.

// Furthermore, because you've been especially nice this year, Santa has mailed you instructions on how to display the ideal lighting configuration.

// Lights in your grid are numbered from 0 to 999 in each direction; the lights at each corner are at 0,0, 0,999, 999,999, and 999,0. The instructions include whether to turn on, turn off, or toggle various inclusive ranges given as coordinate pairs. Each coordinate pair represents opposite corners of a rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore refers to 9 lights in a 3x3 square. The lights all start turned off.

// To defeat your neighbors this year, all you have to do is set up your lights by doing the instructions Santa sent you in order.

// For example:

//     turn on 0,0 through 999,999 would turn on (or leave on) every light.
//     toggle 0,0 through 999,0 would toggle the first line of 1000 lights, turning off the ones that were on, and turning on the ones that were off.
//     turn off 499,499 through 500,500 would turn off (or leave off) the middle four lights.

// After following the instructions, how many lights are lit?

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
func (d Day06) SolvePart1(input []string) (int, error) {
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
	return result, nil
}

// You just finish implementing your winning light pattern when you realize you mistranslated Santa's message from Ancient Nordic Elvish.
//
// The light grid you bought actually has individual brightness controls; each light can have a brightness of zero or more. The lights all start at zero.
//
// The phrase turn on actually means that you should increase the brightness of those lights by 1.
//
// The phrase turn off actually means that you should decrease the brightness of those lights by 1, to a minimum of zero.
//
// The phrase toggle actually means that you should increase the brightness of those lights by 2.
//
// What is the total brightness of all lights combined after following Santa's instructions?
//
// For example:
//
//     turn on 0,0 through 0,0 would increase the total brightness by 1.
//     toggle 0,0 through 999,999 would increase the total brightness by 2000000.
//

// SolvePart2 implements solutions.Solver.
func (d Day06) SolvePart2(input []string) (int, error) {
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
	return result, nil
}

func init() {
	solutions.Register(2015, 06, Day06{})
}
