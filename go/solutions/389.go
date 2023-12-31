package solutions

import "strings"

/* Question **********************************
389. Find the Difference
You are given two strings s and t.
String t is generated by random shuffling string s and then add one more letter at a random position.
Return the letter that was added to t.

Example 1:
Input: s = "abcd", t = "abcde"
Output: "e"
Explanation: 'e' is the letter that was added.

Example 2:
Input: s = "", t = "y"
Output: "y"

Constraints:
0 <= s.length <= 1000
t.length == s.length + 1
s and t consist of lowercase English letters.
*********************************************/

func Output389() any {
	return findTheDifference("abcd", "abcde")
}

// * Solution 1
func findTheDifference(s string, t string) byte {
	if len(s) == 0 {
		return byte(t[0])
	}
	for _, v := range s {
		if strings.Contains(s, string(v)) {
			t = strings.Replace(t, string(v), "", 1)
		}
	}
	return byte(t[0])
}
