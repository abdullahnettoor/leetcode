package solutions

/* Question *******************************
392. Is Subsequence
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:
Input: s = "abc", t = "ahbgdc"
Output: true

Example 2:
Input: s = "axc", t = "ahbgdc"
Output: false

Constraints:
0 <= s.length <= 100
0 <= t.length <= 104
s and t consist only of lowercase English letters.
******************************************/

func Output392() any {
	return isSubsequence("abc", "ahbgdc")
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	var k int
	var res string
	for i := range t {
		if k < len(s) && t[i] == s[k] {
			res += string(t[i])
			k += 1
		}
	}
	return res == s
}

