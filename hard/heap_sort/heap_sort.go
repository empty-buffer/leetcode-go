package heapsort

func HeapSort(nums []int) {
	n := len(nums)
	buildMaxHeap(nums)

	for i := n - 1; i > 0; i-- {
		// Swap max element (nums[0]) with the last element nums[i]
		nums[0], nums[i] = nums[i], nums[0]

		// Heapify root element for reduced heap
		heapify(0, i, nums)
	}
}

func buildMaxHeap(h []int) {
	n := len(h)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(i, n, h)
	}
}

func heapify(i, n int, h []int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && h[left] > h[largest] {
		largest = left
	}

	if right < n && h[right] > h[largest] {
		largest = right
	}

	if largest != i {
		h[i], h[largest] = h[largest], h[i]
		heapify(largest, n, h)
	}
}
