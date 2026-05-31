package graphs

type Grap interface {
	AddUndirectedEdge(from string, to string, cost int)
	NodeCount() int
}
