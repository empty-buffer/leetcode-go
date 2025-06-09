package parentheses

/*
20. Valid Parentheses

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
Input: s = "()[]{}"
Output: true

*/

func parenthesesV1(s string) bool {
	var order []rune

	for _, b := range s {
		switch b {
		case '(':
			order = append(order, ')')
		case '[':
			order = append(order, ']')
		case '{':
			order = append(order, '}')
		default:
			if len(order) > 0 && order[len(order)-1] == b {
				order = order[:len(order)-1]
			} else {
				return false
			}
		}
	}

	if len(order) != 0 {
		return false
	}

	return true
}
