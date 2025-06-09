package evaluateexpressiontree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Evaluate_Tree(t *testing.T) {
	tree := New()

	tree.root.Value = -1

	tree.root.Left = &Node{Value: -2}
	tree.root.Right = &Node{Value: -3}

	tree.root.Left.Left = &Node{Value: -4}
	tree.root.Left.Right = &Node{Value: 2}
	tree.root.Right.Left = &Node{Value: 8}
	tree.root.Right.Right = &Node{Value: 3}

	tree.root.Left.Left.Left = &Node{Value: 2}
	tree.root.Left.Left.Right = &Node{Value: 3}

	assert.Equal(t, 7, Evaluate(tree.root))
}
