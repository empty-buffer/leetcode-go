package palindrome

import "fmt"

func palindromeV1(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		// negative numbers are not palindromes
		// also if the last digit of the number is 0, in order to be a palindrome,
		// the number itself should be zero
		return false
	}

	revertedNumber := 0

	// revert the second half of the number and compare with the first half
	for x > revertedNumber {
		fmt.Println("Num", x)
		fmt.Println("step_1", revertedNumber*10)
		fmt.Println("step_2", x%10)

		revertedNumber = revertedNumber*10 + x%10
		fmt.Println(revertedNumber)
		x /= 10
		fmt.Println(x)
	}

	// when the length is an odd number, we can get rid of the middle digit by revertedNumber/10
	return x == revertedNumber || x == revertedNumber/10
}
