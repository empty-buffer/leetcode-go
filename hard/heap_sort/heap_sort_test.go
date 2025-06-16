package heapsort_test

import (
	"testing"

	heapsort "github.com/empty-buffer/leetcode-go/hard/heap_sort"
	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
	}{
		{
			name: "",
			nums: []int{3, 9, 2, 1, 12, 4, 5, 14},
			want: []int{1, 2, 3, 4, 5, 9, 12, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heapsort.HeapSort(tt.nums)
			assert.Equal(t, tt.want, tt.nums)
		})
	}
}
