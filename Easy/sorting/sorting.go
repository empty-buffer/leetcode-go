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

func quickSort(nums []int, low, high int) {
	if low < high {

		pi := partition(nums, low, high)

		quickSort(nums, low, pi-1)
		quickSort(nums, pi+1, high)
	}
}
func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[high] = nums[high], nums[i+1]

	return i + 1
}
