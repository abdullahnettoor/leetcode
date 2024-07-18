package solutions

import "strings"

/* Question **********************************
? 290. Word Pattern
Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s. Specifically:

Each letter in pattern maps to exactly one unique word in s.
Each unique word in s maps to exactly one letter in pattern.
No two letters map to the same word, and no two words map to the same letter.

Example 1:
Input: pattern = "abba", s = "dog cat cat dog"
Output: true

Explanation:
The bijection can be established as:
'a' maps to "dog".
'b' maps to "cat".

Example 2:
Input: pattern = "abba", s = "dog cat cat fish"
Output: false

Example 3:
Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false


Constraints:
1 <= pattern.length <= 300
pattern contains only lower-case English letters.
1 <= s.length <= 3000
s contains only lowercase English letters and spaces ' '.
s does not contain any leading or trailing spaces.
All the words in s are separated by a single space.
*********************************************/

func Output290() any {
	return wordPattern("aaaa", "dog cat cat dog")
}

// * Solution 1 -- Hash Map + Set -- Time: O(n), Space: O(n)
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	m := map[rune]string{}
	added := map[string]bool{}

	for i, char := range pattern {
		key := char

		if val, found := m[key]; !found {

			if added[words[i]] {
				return false
			}
			m[key] = words[i]
			added[words[i]] = true

		} else {
			if val != words[i] {
				return false
			}
		}
	}

	return true
}

// * Solution 2 -- Two Maps -- Time: O(n), Space: O(n)
func wordPattern2(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	patternToWord := make(map[byte]string)
	wordToPattern := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		w := words[i]

		if word, exists := patternToWord[p]; exists {
			if word != w {
				return false
			}
		} else {
			patternToWord[p] = w
		}

		if pat, exists := wordToPattern[w]; exists {
			if pat != p {
				return false
			}
		} else {
			wordToPattern[w] = p
		}
	}

	return true
}


// * Solution 3 -- Single Map with Index -- Time: O(n), Space: O(n)
func wordPattern3(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	indexMap := make(map[any]int)

	for i := 0; i < len(pattern); i++ {
		// Convert pattern char and word to interface{} to store in same map
		p := any(pattern[i])
		w := any(words[i])

		// If first occurrence of either pattern or word,
		// they should both be first occurrences
		if indexMap[p] != indexMap[w] {
			return false
		}

		// Store next index for both
		indexMap[p] = i + 1
		indexMap[w] = i + 1
	}

	return true
}

