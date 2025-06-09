package evaluateexpressiontree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	root Node
}

func New() Tree {
	return Tree{Node{}}
}

func Evaluate(n Node) int {
	if n.Value >= 0 {
		return n.Value
	}

	left := Evaluate(*n.Left)
	right := Evaluate(*n.Right)

	switch n.Value {
	case -1:
		return left + right
	case -2:
		return left - right
	case -3:
		return left / right
	default:
		return left * right
	}
}
