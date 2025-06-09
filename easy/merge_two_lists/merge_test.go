package merge_two_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	listNodeBuilder := func(nums []int) *ListNode {
		if len(nums) == 0 {
			return nil
		}

		head := &ListNode{Val: nums[0]}
		current := head
		for _, num := range nums[1:] {
			current.Next = &ListNode{Val: num}
			current = current.Next
		}
		return head
	}

	listNodeValues := func(node *ListNode) []int {
		var values []int
		for node != nil {
			values = append(values, node.Val)
			node = node.Next
		}
		return values
	}

	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "All elements distinct, in order",
			l1:       []int{1, 3, 5},
			l2:       []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Empty lists",
			l1:       []int{},
			l2:       []int{},
			expected: nil,
		},
		{
			name:     "Mix of positive and negative elements",
			l1:       []int{-3, -1, 3},
			l2:       []int{-2, 2, 4},
			expected: []int{-3, -2, -1, 2, 3, 4},
		},
		{
			name:     "L1 is empty",
			l1:       []int{},
			l2:       []int{2, 4, 6},
			expected: []int{2, 4, 6},
		},
		{
			name:     "L2 is empty",
			l1:       []int{1, 3, 5},
			l2:       []int{},
			expected: []int{1, 3, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l1 := listNodeBuilder(test.l1)
			l2 := listNodeBuilder(test.l2)
			actual := mergeTwoLists(l1, l2)
			assert.Equal(t, test.expected, listNodeValues(actual))
		})
	}
}
