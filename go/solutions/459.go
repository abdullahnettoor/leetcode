package solutions

import "strings"

/* Question **********************************
? 459. Repeated Substring Pattern
Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.

Example 1:
Input: s = "abab"
Output: true
Explanation: It is the substring "ab" twice.

Example 2:
Input: s = "aba"
Output: false

Example 3:
Input: s = "abcabcabcabc"
Output: true
Explanation: It is the substring "abc" four times or the substring "abcabc" twice.

Constraints:
1 <= s.length <= 104
s consists of lowercase English letters.
*********************************************/

func Output459() any {
	return repeatedSubstringPattern("abcab")
}

// * Solution 1
func repeatedSubstringPattern(s string) bool {
	t := s + s
	return strings.Contains(t[1:len(t)-1], s)
}
