package sorting

// ----- Quick Sort With Partition -----

func quickSortWithPartition(nums []int, low, high int) {
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

// ----- Quick Sort -----

func quickSort(nums []int, low, high int) {
	if low >= high {
		return
	}

	pivot := low
	leftIdx := low + 1
	rightIdx := high

	for leftIdx <= rightIdx {
		if nums[leftIdx] > nums[pivot] && nums[high] < nums[pivot] {
			nums[leftIdx], nums[high] = nums[high], nums[leftIdx]
		}

		if nums[leftIdx] <= nums[pivot] {
			leftIdx += 1
		}

		if nums[rightIdx] >= nums[pivot] {
			rightIdx -= 1
		}
	}

	nums[pivot], nums[rightIdx] = nums[rightIdx], nums[pivot]

	leftSmaller := rightIdx-1-leftIdx < high-(rightIdx+1)

	if leftSmaller {
		quickSort(nums, low, rightIdx-1)
		quickSort(nums, rightIdx+1, high)
	} else {
		quickSort(nums, rightIdx+1, high)
		quickSort(nums, low, rightIdx+1)
	}
}
