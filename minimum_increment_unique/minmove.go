package minimum_increment_unique

import (
	"fmt"
	"slices"
)

/*945. Minimum Increment to Make Array Unique
Medium
Topics
Companies
You are given an integer array nums. In one move, you can pick an index i where 0 <= i < nums.length and increment nums[i] by 1.

Return the minimum number of moves to make every value in nums unique.

The test cases are generated so that the answer fits in a 32-bit integer.

The notation '1e8' in visuals in typically represents values using scientific notation. In this case, it means 1 multiplied by 10 to the power of 8, which is 100,000,000

Example 1:

Input: nums = [1,2,2]
Output: 1
Explanation: After 1 move, the array could be [1, 2, 3].

Example 2:

Input: nums = [3,2,1,2,1,7]
Output: 6
Explanation: After 6 moves, the array could be [3, 4, 1, 2, 5, 7].
It can be shown with 5 or less moves that it is impossible for the array to have all unique values.
*/

func minMoveV1(numbers []int) int {
	slices.Sort(numbers)

	fmt.Println(numbers)

	move := 0

	for i := 1; i < len(numbers); i++ {
		// If the current number is less than or equal to the one in the previous index
		if numbers[i] <= numbers[i-1] {
			// Increment the 'move' counter by the difference between the current and the previous number plus one
			move += numbers[i-1] - numbers[i] + 1
			// Update numbers[i] to be one greater than the previous number
			numbers[i] = numbers[i-1] + 1
		}
	}

	fmt.Println(numbers)
	return move
}

func minMoveV2(numbers []int) int {
	const maxNum = 1e+2
	if len(numbers) < 2 {
		return 0
	}

	res := 0

	// Creating matrix of numbers what were found in the slice
	count := make([]int, maxNum)
	for _, v := range numbers {
		count[v]++
	}

	// Calculating number of moves
	for i, v := range count {
		if v > 1 {
			res += count[i] - 1
			count[i+1] += count[i] - 1
		}
	}

	return res
}
