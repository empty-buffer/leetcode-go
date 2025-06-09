package search_insert_position

/*
Problem condition
35. Search Insert Position

Given a sorted array of distinct integers and a target value,
return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4
*/

func searchInsertPositionV1(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums) / 2

	for target < nums[n] {
		n = n / 2
	}

	for i, num := range nums[n:] {
		if num == target {
			return n + i
		}

		if num > target {
			return n + i
		}
	}

	return len(nums)
}

func searchInsertPositionV2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
