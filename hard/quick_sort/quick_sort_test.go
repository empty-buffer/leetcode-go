package sorting

import (
	"fmt"
	"testing"
)

func Test_QuickSort(t *testing.T) {
	nums := []int{5, 2, 6, 3, 1, 5, 4}
	fmt.Println("Before sorting:", nums)

	quickSort(nums, 0, len(nums)-1)

	fmt.Println("After sorting:", nums)
}

func Test_QuickSortWithPartittion(t *testing.T) {
	nums := []int{5, 2, 6, 3, 1, 5, 4}
	fmt.Println("Before sorting:", nums)

	quickSortWithPartition(nums, 0, len(nums)-1)

	fmt.Println("After sorting:", nums)
}
