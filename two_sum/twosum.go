package two_sum

/*
1. Two Sum
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

func twoSumV1(nums []int, target int) []int {
	for i := range nums {
		for index := range nums {
			if index != i {
				if nums[index]+nums[i] == target {
					return []int{i, index}
				}
			}
		}
	}
	return []int{}
}

func twoSumV2(nums []int, target int) []int {
	// [number]index
	targetMap := make(map[int]int)
	for i := range nums {
		if t, ok := targetMap[target-nums[i]]; ok {
			return []int{t, i}
		} else {
			targetMap[nums[i]] = i
		}
	}
	return []int{}
}
