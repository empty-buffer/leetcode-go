package majority

import (
	"fmt"
	"slices"
)

// Description
//
//Given an array nums of size n, return the majority element.
//
//The majority element is the element that appears more than
//⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
//
//Input: nums = [2,2,1,1,1,2,2]

// Complexity
// Time complexity:
// O(n)O(n)O(n)
//
// Space complexity:
// Approach: Boyer-Moore Majority Vote Algorithm (O(1)O(1)O(1)

func findMajorityElementV1(numbers []int) int {
	var (
		majorityIndex     = 0
		majorityThreshold = 0
		processedNumbers  = make(map[int]bool)
	)

	for i := 0; i < len(numbers); i++ {
		if processedNumbers[i] {
			continue
		}

		occurrenceCount := 0
		for _, n := range numbers {
			if n == numbers[i] {
				occurrenceCount++
			}
		}
		//occurrenceCount := countOccurrence(numbers, numbers[currentNumber])

		if occurrenceCount > majorityThreshold {
			majorityThreshold = occurrenceCount
			majorityIndex = i
		}
		processedNumbers[numbers[i]] = true
	}

	return numbers[majorityIndex]
}

func findMajorityElementV2(numbers []int) int {
	var (
		majorityIndex     = 0
		majorityThreshold = 0
		processedNumbers  = make(map[int]bool)
	)

	for index := 0; index < len(numbers); index++ {
		currentNumber := numbers[index]

		if processedNumbers[currentNumber] {
			continue
		}

		occurrenceCount := countOccurrence(numbers, currentNumber)

		if occurrenceCount > majorityThreshold {
			majorityIndex = index
			majorityThreshold = occurrenceCount
		}

		processedNumbers[currentNumber] = true
	}

	return numbers[majorityIndex]
}

func countOccurrence(numbers []int, target int) int {
	count := 0

	for _, number := range numbers {
		if number == target {
			count++
		}
	}

	return count
}

func findMajorityElementV3(nums []int) int {
	majorityElement, majorityElementFrequency := nums[0], 1
	for i := 0; i < len(nums); i++ {
		if majorityElementFrequency == 0 {
			majorityElement, majorityElementFrequency = nums[i], 1
		} else {
			if nums[i] == majorityElement {
				majorityElementFrequency += 1
			} else {
				majorityElementFrequency -= 1
			}
		}
	}
	return majorityElement
}

func findMajorityElementV4(nums []int) int {
	slices.Sort(nums)
	fmt.Println(len(nums) / 2)

	count := 0
	for _, num := range nums {
		tempIndex := num
		if tempIndex == num {
			count++
			if count > len(nums)/2 {
				return tempIndex
			}
			continue
		}
		tempIndex = num
		count = 0
	}
	return 0
}
