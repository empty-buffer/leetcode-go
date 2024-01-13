package stack_sequences

/*Given two integer arrays pushed and popped each with distinct values, return true
if this could have been the result of a sequence of push and pop operations on an initially empty stack, or false otherwise.

Example 1:

Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
Output: true
Explanation: We might do the following sequence:
push(1), push(2), push(3), push(4),
pop() -> 4,
push(5),
pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

Example 2:

Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
Output: false
Explanation: 1 cannot be popped before 2.*/

func sequencesV2(pushed []int, popped []int) bool {
	var stack []int
	n := len(pushed)
	var p1, p2 int

	for p1 < n {
		stack = append(stack, pushed[p1])
		p1++
		for p2 < n && len(stack) > 0 && popped[p2] == stack[len(stack)-1] {
			p2++
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func sequencesV1(pushed, popped []int) bool {
	var (
		stack    = make([]int, 0, len(pushed))
		popIndex int
	)

	for _, value := range pushed {
		stack = push(stack, value)

		for len(stack) > 0 {
			if stack[len(stack)-1] == popped[popIndex] {
				stack = pop(stack)
				popIndex++
			} else {
				break
			}
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == popped[popIndex] {
			stack = pop(stack)

		} else {
			return false
		}
		popIndex++
	}

	return true
}

func push(slice []int, value int) []int {
	return append(slice, value)
}

func pop(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	return slice[:len(slice)-1]
}
