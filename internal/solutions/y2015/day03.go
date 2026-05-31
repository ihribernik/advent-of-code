package y2015

import (
	"slices"
	"strings"

	"github.com/ihribernik/aoc-cli/internal/directions"
)

type Day03 struct{}

func (d Day03) parseLine(input []string) []string {
	inputVars := strings.Split(input[0], "")
	return inputVars
}

func (d Day03) SolvePart1(input []string) (int, error) {
	position := directions.NewDirection(0, 0)

	visited := []directions.Direction{
		position,
	}

	characters := d.parseLine(input)

	for _, direction := range characters {
		currentDirection := directions.DIRECTIONS[direction]
		position = currentDirection.NewPositionWith(position)
		if !slices.Contains(visited, position) {
			visited = append(visited, position)
		}
	}
	return len(visited), nil
}

func (d Day03) SolvePart2(input []string) (int, error) {
	position := directions.NewDirection(0, 0)

	positions := []directions.Direction{
		position,
		position,
	}

	visited := []directions.Direction{
		position,
	}

	inputDirections := d.parseLine(input)

	for chunk := range slices.Chunk(inputDirections, 2) {
		for i, direction := range chunk {
			currentDirection := directions.DIRECTIONS[direction]
			positions[i] = currentDirection.NewPositionWith(positions[i])
			if !slices.Contains(visited, positions[i]) {
				visited = append(visited, positions[i])
			}
		}
	}
	return len(visited), nil
}
