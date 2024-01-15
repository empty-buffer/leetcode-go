package rotate_array

/* Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]

Constraints:

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105

Follow up:
Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
Could you do it in-place with O(1) extra space? */

func rotateV1(numbers []int, steps int) []int {
	var (
		index = len(numbers) - 1
	)

	for ; steps > 0; steps-- {
		moving := numbers[index]
		copy(numbers[1:], numbers[:index+1])
		numbers[0] = moving
	}

	return numbers
}

func rotateV2(numbers []int, steps int) []int {
	if steps > len(numbers) {
		steps = steps % len(numbers)
	}
	index := len(numbers) - steps
	_nums := append(numbers[index:], numbers[0:index]...)
	for i := 0; i < len(_nums); i++ {
		numbers[i] = _nums[i]
	}
	return numbers
}

func rotateV3(numbers []int, steps int) {
	if steps > len(numbers) {
		steps = steps % len(numbers)
	}

	if steps == 0 {
		return
	}

	_numbers := append(numbers[len(numbers)-steps:], numbers[0:len(numbers)-steps]...)
	copy(numbers, _numbers)
}

// rotateV4 rotates the elements in the numbers slice by the given number of steps.
// It modifies the numbers slice in-place.
//
// If the number of steps is greater than the length of the numbers slice,
// the steps are limited to the length of the numbers slice.
//
// This function uses the Reverse function to reverse sections of the numbers slice
// in order to achieve the rotation.
func rotateV4(numbers []int, steps int) {
	steps %= len(numbers) // limit steps to length of numbers
	println(steps)

	reverse(numbers)
	reverse(numbers[:steps])
	reverse(numbers[steps:])
}

func reverse(slice []int) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}
