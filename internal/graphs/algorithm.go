package graphs

import "math"

type pathSearchState struct {
	Visited map[string]bool
	Best    int
}

func ShortestHamiltonianPath(g Graph) int {
	edges := g.GetEdges()

	if g == nil || len(edges) == 0 {
		return 0
	}

	state := &pathSearchState{
		Visited: make(map[string]bool, len(edges)),
		Best:    math.MaxInt,
	}

	for start := range edges {
		state.Visited[start] = true
		walkShortest(g, state, start, 1, 0)
		state.Visited[start] = false
	}

	return state.Best
}

func walkShortest(g Graph, state *pathSearchState, node string, count int, total int) {
	if total >= state.Best {
		return
	}

	edges := g.GetEdges()

	if count == len(edges) {
		state.Best = total
		return
	}

	for _, edge := range edges[node] {
		if state.Visited[edge.To] {
			continue
		}

		state.Visited[edge.To] = true
		walkShortest(g, state, edge.To, count+1, total+edge.Cost)
		state.Visited[edge.To] = false
	}
}

func LongestHamiltonianPath(g Graph) int {
	edges := g.GetEdges()

	if g == nil || len(edges) == 0 {
		return 0
	}

	state := &pathSearchState{
		Visited: make(map[string]bool, len(edges)),
	}

	for start := range edges {
		state.Visited[start] = true
		walkLongest(g, state, start, 1, 0)
		state.Visited[start] = false
	}

	return state.Best
}

func walkLongest(g Graph, state *pathSearchState, node string, count int, total int) {
	if count == len(g.GetEdges()) {
		if total > state.Best {
			state.Best = total
		}
		return
	}
	edges := g.GetEdges()
	for _, edge := range edges[node] {
		if state.Visited[edge.To] {
			continue
		}

		state.Visited[edge.To] = true
		walkLongest(g, state, edge.To, count+1, total+edge.Cost)
		state.Visited[edge.To] = false
	}
}
