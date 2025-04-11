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

	visited := map[string]bool{
		position.String(): true,
	}

	characters := d.parseLine(input)

	for _, direction := range characters {
		position = common.DIRECTIONS[direction].NewPositionWith(position)
		visited[position.String()] = true
	}

	return len(visited), nil
}

func (d Day03) SolvePart2(input []string) (int, error) {
	position := common.Direction{X: 0, Y: 0}

	visited := map[string]bool{
		position.String(): true,
	}
	directions := strings.Split(input[0], "")

	for chunk := range slices.Chunk(directions, 2) {
		for _, direction := range chunk {
			println(direction)
			position := common.DIRECTIONS[direction].NewPosition(position, 1)
			visited[position.String()] = true
		}

	}
	return len(visited), nil
}

func init() {
	solutions.Register(2015, 03, Day03{})
}
