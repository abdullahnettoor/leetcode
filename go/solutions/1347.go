package solutions

/* Question **********************************
? 1347. Minimum Number of Steps to Make Two Strings Anagram
You are given two strings of the same length s and t. In one step you can choose any character of t and replace it with another character.
Return the minimum number of steps to make t an anagram of s.
An Anagram of a string is a string that contains the same characters with a different (or the same) ordering.

Example 1:
Input: s = "bab", t = "aba"
Output: 1
Explanation: Replace the first 'a' in t with b, t = "bba" which is anagram of s.

Example 2:
Input: s = "leetcode", t = "practice"
Output: 5
Explanation: Replace 'p', 'r', 'a', 'i' and 'c' from t with proper characters to make t anagram of s.

Example 3:
Input: s = "anagram", t = "mangaar"
Output: 0
Explanation: "anagram" and "mangaar" are anagrams.

Constraints:
1 <= s.length <= 5 * 104
s.length == t.length
s and t consist of lowercase English letters only.
*********************************************/

func Output1347() any {
	return minSteps("ander", "enter")
}

// * Solution 1 -- Using Hash Tables (i.e; map)
func minSteps(s string, t string) int {
	steps1, steps2 := 0, 0
	m1, m2 := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s); i++ {
		m1[s[i]]++
		m2[t[i]]++
	}
	for k := range m1 {
		if m1[k] > m2[k] {
			steps1 += m1[k] - m2[k]
		}
	}
	for k := range m2 {
		if m2[k] > m1[k] {
			steps2 += m2[k] - m1[k]
		}
	}
	return min(steps1, steps2)
}

// TODO: * Solution 2 -- Using Array
