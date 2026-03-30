package nodes

type Node struct {
	Src      *Node
	Dst      *Node
	Distance int
}

func NewNode(distance int) *Node {
	return &Node{Src: nil, Dst: nil, Distance: distance}
}
