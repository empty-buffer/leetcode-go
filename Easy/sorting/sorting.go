package sorting

func bubbleSort(nums []int) {
	n := len(nums)

	for i := range n - 1 {
		swapped := false

		for j := range n - 1 - i {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
