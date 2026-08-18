package y2015

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ihribernik/aoc-cli/internal/graphs"
)

type Day09 struct{}

func parseDay09Graph(input []string) (graphs.Graph, error) {
	graph := graphs.New(len(input))

	for i, line := range input {
		parts := strings.Fields(line)
		if len(parts) != 5 || parts[1] != "to" || parts[3] != "=" || parts[0] == "" || parts[2] == "" {
			return nil, newInputError(9, i+1, line, errors.New("expected source to destination = distance"))
		}
		src, dst, dist := parts[0], parts[2], parts[4]
		distInt, err := strconv.Atoi(dist)
		if err != nil {
			return nil, newInputError(9, i+1, line, fmt.Errorf("invalid distance %q: %w", dist, err))
		}
		if distInt < 0 {
			return nil, newInputError(9, i+1, line, errors.New("distance must not be negative"))
		}

		graph.AddUndirectedEdge(src, dst, distInt)
	}

	return graph, nil
}

func (d Day09) SolvePart1(input []string) (int, error) {
	graph, err := parseDay09Graph(input)
	if err != nil {
		return 0, err
	}

	return graphs.ShortestHamiltonianPath(graph), nil
}

func (d Day09) SolvePart2(input []string) (int, error) {
	graph, err := parseDay09Graph(input)
	if err != nil {
		return 0, err
	}

	return graphs.LongestHamiltonianPath(graph), nil
}
