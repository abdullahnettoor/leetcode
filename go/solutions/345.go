package solutions

/* Question **********************************
? 345. Reverse Vowels of a String
Given a string s, reverse only all the vowels in the string and return it.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

Example 1:
Input: s = "hello"
Output: "holle"

Example 2:
Input: s = "leetcode"
Output: "leotcede"

Constraints:
1 <= s.length <= 3 * 105
s consist of printable ASCII characters.
*********************************************/

func Output345() any {
	return reverseVowels("leetcode")
}

// * Solution 1
func reverseVowels(s string) string {
	c := []byte(s)
	l, r := 0, len(s)-1
	for l <= r {
		if isVowel(c[l]) && isVowel(c[r]) {
			c[l], c[r] = c[r], c[l]
			l++
			r--
			continue
		} else if isVowel(c[l]) && !isVowel(c[r]) {
			r--
		} else {
			l++
		}
	}
	return string(c)
}

func isVowel(v byte) bool {
	switch rune(v) {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
