package solutions

/* Question ********************************
? 242. Valid Anagram
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Constraints:
1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.
*******************************************/

func Output242() any {
	return isAnagram("anagram", "nagaram")
}

// * Solution 1
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int, 26)
	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}

// * Solution 2
// * The following solution is not efficent. Will be update with new solution later.
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	sChars := strings.Split(s, "")
// 	tChars := strings.Split(t, "")
// 	sort.Strings(sChars)
// 	sort.Strings(tChars)

// 	return reflect.DeepEqual(sChars, tChars)
// }
