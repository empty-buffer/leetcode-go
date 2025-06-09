package common_prefix

import (
	"sort"
)

/*
14. Longest Common Prefix
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
*/

func longestCommonPrefixV1(strs []string) string {
	var minLength int
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	minLength = len(strs[0])

	var prefixResult string

	for i := 0; i < minLength; i++ {
		p := strs[0][i]
		for _, s := range strs {
			if s[i] == p {
				continue
			} else {
				return prefixResult
			}
		}
		prefixResult = strs[0][:i+1]
	}

	return prefixResult
}

// longestCommonPrefixV2 finds the longest common prefix among a slice of strings.
// It returns an empty string if the slice is empty.
// The function iterates through each string in the slice, comparing each character
// at the same index. It stops when it encounters a difference or reaches the end of
// the shortest string. The prefix is then updated to include only the common
// characters. This process is repeated for all strings. The final result is the
// longest common prefix.
func longestCommonPrefixV2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	p := strs[0]
	for _, s := range strs {
		i := 0
		for ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {
		}
		p = p[:i]
	}

	return p
}

func longestCommonPrefixV3(strs []string) string {
	if len(strs) == 1 { // handle only 1 element
		return strs[0]
	}

	if len(strs) == 0 { // handle empty slice
		return ""
	}

	// sort them first, the most different one will be in first and last
	sort.Strings(strs)

	// compare first and last
	l := len(strs)
	for i := range strs[0] {
		if strs[0][i] != strs[l-1][i] {
			return strs[0][:i]
		}
	}
	return strs[0]
}
