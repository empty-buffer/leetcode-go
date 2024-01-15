package roman_to_integer

/*
13. Roman to Integer
Easy

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.
*/

var convertor = make(map[rune][]int)

func initConvertor() {
	convertor['I'] = []int{0, 1}
	convertor['V'] = []int{1, 5}
	convertor['X'] = []int{2, 10}
	convertor['L'] = []int{3, 50}
	convertor['C'] = []int{4, 100}
	convertor['D'] = []int{5, 500}
	convertor['M'] = []int{6, 1000}
}

func romanToIntV1(romanNumber string) int {
	initConvertor()
	result := 0

	for i := 0; i < len(romanNumber); i++ {
		if value, ok := convertor[rune(romanNumber[i])]; ok {
			if i+1 < len(romanNumber) {
				nextRM := getIndex(rune(romanNumber[i+1]))
				if nextRM[0] > value[0] {
					result += nextRM[1] - value[1]
					i++
				} else {
					result += value[1]
				}
			} else {
				result += value[1]
			}
		} else {
			return 0
		}
	}
	return result
}

func getIndex(r rune) []int {
	if value, ok := convertor[r]; ok {
		return value
	}
	return []int{0, 0}
}

func romanToIntV2(s string) int {
	sum := 0

	rim := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i, v := range s {
		// add each number to sum

		sum += rim[v]
		if i != 0 {
			// If previous number less than current one, multiple previous num by 2 and substract from sum.
			if rim[rune(s[i-1])] < rim[(v)] {
				sum -= 2 * rim[rune(s[i-1])]
			}
		}
	}

	return sum
}
