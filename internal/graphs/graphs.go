package graphs

import "math"

type Edge struct {
	To   string
	Cost int
}

type Graph struct {
	Edges map[string][]Edge
}

func New(size int) *Graph {
	return &Graph{
		Edges: make(map[string][]Edge, size),
	}
}

func (g *Graph) AddUndirectedEdge(from string, to string, cost int) {
	g.Edges[from] = append(g.Edges[from], Edge{To: to, Cost: cost})
	g.Edges[to] = append(g.Edges[to], Edge{To: from, Cost: cost})
}

func (g *Graph) NodeCount() int {
	return len(g.Edges)
}

type pathSearchState struct {
	Visited map[string]bool
	Best    int
}

func ShortestHamiltonianPath(g *Graph) int {
	if g == nil || len(g.Edges) == 0 {
		return 0
	}

	state := &pathSearchState{
		Visited: make(map[string]bool, len(g.Edges)),
		Best:    math.MaxInt,
	}

	for start := range g.Edges {
		state.Visited[start] = true
		walkShortest(g, state, start, 1, 0)
		state.Visited[start] = false
	}

	return state.Best
}

func walkShortest(g *Graph, state *pathSearchState, node string, count int, total int) {
	if total >= state.Best {
		return
	}
	if count == len(g.Edges) {
		state.Best = total
		return
	}

	for _, edge := range g.Edges[node] {
		if state.Visited[edge.To] {
			continue
		}

		state.Visited[edge.To] = true
		walkShortest(g, state, edge.To, count+1, total+edge.Cost)
		state.Visited[edge.To] = false
	}
}

func LongestHamiltonianPath(g *Graph) int {
	if g == nil || len(g.Edges) == 0 {
		return 0
	}

	state := &pathSearchState{
		Visited: make(map[string]bool, len(g.Edges)),
	}

	for start := range g.Edges {
		state.Visited[start] = true
		walkLongest(g, state, start, 1, 0)
		state.Visited[start] = false
	}

	return state.Best
}

func walkLongest(g *Graph, state *pathSearchState, node string, count int, total int) {
	if count == len(g.Edges) {
		if total > state.Best {
			state.Best = total
		}
		return
	}

	for _, edge := range g.Edges[node] {
		if state.Visited[edge.To] {
			continue
		}

		state.Visited[edge.To] = true
		walkLongest(g, state, edge.To, count+1, total+edge.Cost)
		state.Visited[edge.To] = false
	}
}
