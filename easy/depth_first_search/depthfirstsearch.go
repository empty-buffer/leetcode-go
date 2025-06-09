package depthfirstsearch

type Node struct {
	Value string
	Nodes []Node
}

type Constractor interface {
	Build() Node
	AddChildren(nodes ...Node) Constractor
}

func NewNode(v string) Constractor {
	return &Node{
		Value: v,
		Nodes: []Node{},
	}
}

func (n *Node) AddChildren(nodes ...Node) Constractor {
	n.Nodes = append(n.Nodes, nodes...)

	return n
}

func (n *Node) Build() Node {
	return *n
}

func DepthFirstSearch(res []string, n Node) []string {
	if n.Value == "" {
		return res
	}

	res = append(res, n.Value)

	for _, cn := range n.Nodes {
		res = DepthFirstSearch(res, cn)
	}

	return res
}
