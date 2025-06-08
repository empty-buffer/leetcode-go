package minimumwaitingtime

func bubbleSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := range len(nums) - 1 {
		for j := range len(nums) - 1 - i {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

func MinimumWaitingTime(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	sorted := bubbleSort(nums)

	total := 0

	for i := range len(nums) - 1 {
		total += sorted[i] * (len(sorted) - 1 - i)

	}

	return total
}
