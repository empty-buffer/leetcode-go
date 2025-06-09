package first_occurrence

/*
Problem condition

Code
28. Find the Index of the First Occurrence in a String

Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
*/
// sadbutsad sad
func firstOccurrenceV1(haystack string, needle string) int {
	if haystack == "" {
		return -1
	}

	var n = 0

	for h, c := range haystack {
		if string(c) == string(needle[n]) {
			for ; n < len(needle) && h+n < len(haystack) && needle[n] == haystack[h+n]; n++ {
				if n == len(needle)-1 {
					return h
				}
			}
		}
		n = 0
	}

	return -1
}
