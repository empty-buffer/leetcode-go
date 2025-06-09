package sorting

import (
	"fmt"
	"testing"
)

func Test_BobbleSort(t *testing.T) {
	nums := []int{5, 2, 6, 3, 1, 4}
	fmt.Println("Before sorting:", nums)

	bubbleSort(nums)

	fmt.Println("After sorting:", nums)
}
