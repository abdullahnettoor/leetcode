package solutions

import "strings"

// 14. Longest Common Prefix
// Solved
// Easy
// Topics
// Companies
// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"

// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
 

// Constraints:
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters.

func LongestCommonPrefix(s []string) string {
	var prefix = s[0]
	for i := 1; i < len(s); i++ {
		for !strings.HasPrefix(s[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
