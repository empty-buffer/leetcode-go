package node_depths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Node_Depths(t *testing.T) {
	tree := New()

	tree.Insert(22) // 0
	tree.Insert(10) // 1
	tree.Insert(25) // 1
	tree.Insert(19) // 2
	tree.Insert(9)  // 2
	tree.Insert(4)  // 3
	tree.Insert(50) // 2
	tree.Insert(1)  // 4

	assert.Equal(t, 15, TotalDepth(0, tree.root))
}

func Test_Max_Depth(t *testing.T) {
	tree := New()

	tree.Insert(22) // 0
	tree.Insert(10) // 1
	tree.Insert(25) // 1
	tree.Insert(19) // 2
	tree.Insert(9)  // 2
	tree.Insert(4)  // 3
	tree.Insert(50) // 2
	tree.Insert(1)  // 4

	assert.Equal(t, 4, MaxDepth(0, tree.root))
}

func Test_Max_DepthV1(t *testing.T) {
	tree := New()

	tree.Insert(22) // 0
	tree.Insert(10) // 1
	tree.Insert(25) // 1
	tree.Insert(19) // 2
	tree.Insert(9)  // 2
	tree.Insert(4)  // 3
	tree.Insert(50) // 2
	tree.Insert(1)  // 4

	assert.Equal(t, 4, MaxDepthV2(tree.root))
}
