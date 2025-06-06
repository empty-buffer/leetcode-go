package node_depths

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value int) {
	if n.Value < value {
		if n.Left != nil {
			n.Left.Insert(value)
		} else {
			n.Left = &Node{
				Value: value,
			}
		}
	} else {
		if n.Right != nil {
			n.Right.Insert(value)
		} else {
			n.Right = &Node{
				Value: value,
			}
		}
	}
}

func (n *Node) Travers(name string) {
	if n.Left != nil {
		n.Left.Travers("left")
	}

	if n.Right != nil {
		n.Right.Travers("right")
	}
}

type Tree struct {
	root *Node
}

func New() Tree {
	return Tree{&Node{}}
}

func TotalDepth(depth int, n *Node) int {
	if n == nil {
		return 0
	}

	return depth + TotalDepth(depth+1, n.Left) + TotalDepth(depth+1, n.Right)
}

func MaxDepth(current int, n *Node) int {
	if n == nil {
		return current - 1
	}

	l := MaxDepth(current+1, n.Left)
	r := MaxDepth(current+1, n.Right)

	if l > r {
		return l
	}

	return r
}

func MaxDepthV2(n *Node) int {
	if n == nil {
		return 0
	}

	l := MaxDepthV2(n.Left)
	r := MaxDepthV2(n.Right)

	if l > r {
		return l + 1
	}

	return r + 1
}

func (t *Tree) Insert(value int) {
	if t.root.Value != 0 {
		t.root.Insert(value)
		return
	}

	t.root.Value = value
}

func (t *Tree) Travers() {
	t.root.Travers("root")
}

func NodeDepths() {
	tree := New()

	tree.Insert(22)
	tree.Insert(10)
	tree.Insert(25)
	tree.Insert(19)
	tree.Insert(9)
	tree.Insert(4)

	tree.Travers()

	fmt.Println(TotalDepth(0, tree.root))
}
