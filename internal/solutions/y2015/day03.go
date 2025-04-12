package y2015

import (
	"slices"
	"strings"

	"github.com/ihribernik/advent-of-code/internal/solutions"
	"github.com/ihribernik/advent-of-code/pkg/common"
)

type Day03 struct{}

func (d Day03) parseLine(input []string) []string {
	inputVars := strings.Split(input[0], "")
	return inputVars
}

func (d Day03) SolvePart1(input []string) (int, error) {
	position := common.Direction{X: 0, Y: 0}

	visited := []common.Direction{
		position,
	}

	characters := d.parseLine(input)

	for _, direction := range characters {
		currentDirection := common.DIRECTIONS[direction]
		position = currentDirection.NewPositionWith(position)
		if !slices.Contains(visited, position) {
			visited = append(visited, position)
		}
	}

	return len(visited), nil
}

func (d Day03) SolvePart2(input []string) (int, error) {
	position := common.Direction{X: 0, Y: 0}

	positions := []common.Direction{
		position,
		position,
	}

	visited := []common.Direction{
		position,
	}

	directions := d.parseLine(input)

	for chunk := range slices.Chunk(directions, 2) {
		for i, direction := range chunk {
			currentDirection := common.DIRECTIONS[direction]
			positions[i] = currentDirection.NewPositionWith(positions[i])
			if !slices.Contains(visited, positions[i]) {
				visited = append(visited, positions[i])
			}
		}

	}
	return len(visited), nil
}

func init() {
	solutions.Register(2015, 03, Day03{})
}
