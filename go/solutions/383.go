package solutions

/* Question **********************************
? 383. Ransom Note
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
Each letter in magazine can only be used once in ransomNote.

Example 1:
Input: ransomNote = "a", magazine = "b"
Output: false

Example 2:
Input: ransomNote = "aa", magazine = "ab"
Output: false

Example 3:
Input: ransomNote = "aa", magazine = "aab"
Output: true

Constraints:
1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.
*********************************************/

func Output383() any {
	return canConstruct("aa", "aab")
}

// * Solution 1 -- Using Single loop
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, r := range ransomNote {
		m[r]++
	}
	for _, r := range magazine {
		if _, found := m[r]; found {
			m[r]--
		}
	}
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}
