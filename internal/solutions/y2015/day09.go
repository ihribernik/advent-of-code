package y2015

import (
	"strconv"
	"strings"

	"github.com/ihribernik/aoc-cli/internal/graphs"
)

type Day09 struct{}

func parseDay09Graph(input []string) (*graphs.Graph, error) {
	graph := graphs.New(len(input))

	for _, line := range input {
		parts := strings.Split(line, " ")
		src, dst, dist := parts[0], parts[2], parts[4]
		distInt, err := strconv.Atoi(dist)
		if err != nil {
			return nil, err
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
