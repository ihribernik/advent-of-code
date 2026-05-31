package graphs

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *Graph
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_AddUndirectedEdge(t *testing.T) {
	type args struct {
		from string
		to   string
		cost int
	}
	tests := []struct {
		name string
		g    *Graph
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.AddUndirectedEdge(tt.args.from, tt.args.to, tt.args.cost)
		})
	}
}

func TestGraph_NodeCount(t *testing.T) {
	tests := []struct {
		name string
		g    *Graph
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.NodeCount(); got != tt.want {
				t.Errorf("Graph.NodeCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortestHamiltonianPath(t *testing.T) {
	type args struct {
		g *Graph
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortestHamiltonianPath(tt.args.g); got != tt.want {
				t.Errorf("ShortestHamiltonianPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_walkShortest(t *testing.T) {
	type args struct {
		g     *Graph
		state *pathSearchState
		node  string
		count int
		total int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			walkShortest(tt.args.g, tt.args.state, tt.args.node, tt.args.count, tt.args.total)
		})
	}
}

func TestLongestHamiltonianPath(t *testing.T) {
	type args struct {
		g *Graph
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestHamiltonianPath(tt.args.g); got != tt.want {
				t.Errorf("LongestHamiltonianPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_walkLongest(t *testing.T) {
	type args struct {
		g     *Graph
		state *pathSearchState
		node  string
		count int
		total int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			walkLongest(tt.args.g, tt.args.state, tt.args.node, tt.args.count, tt.args.total)
		})
	}
}
