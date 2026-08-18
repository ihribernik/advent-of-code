package graphs

type Graph interface {
	AddUndirectedEdge(from string, to string, cost int)
	NodeCount() int
	GetEdges() map[string][]edge
}
