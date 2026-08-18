package graphs

type edge struct {
	To   string
	Cost int
}

type graph struct {
	edges map[string][]edge
}

func New(size int) Graph {
	return &graph{
		edges: make(map[string][]edge, size),
	}
}

func (g *graph) GetEdges() map[string][]edge {
	return g.edges
}

func (g *graph) AddUndirectedEdge(from string, to string, cost int) {
	g.edges[from] = append(g.edges[from], edge{To: to, Cost: cost})
	g.edges[to] = append(g.edges[to], edge{To: from, Cost: cost})
}

func (g *graph) NodeCount() int {
	return len(g.edges)
}
