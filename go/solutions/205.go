package solutions

/* Question **********************************
? 205. Isomorphic Strings
Given two strings s and t, determine if they are isomorphic.
Two strings s and t are isomorphic if the characters in s can be replaced to get t.
All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

Example 1:
Input: s = "egg", t = "add"
Output: true

Explanation:
The strings s and t can be made identical by:
Mapping 'e' to 'a'.
Mapping 'g' to 'd'.

Example 2:
Input: s = "foo", t = "bar"
Output: false

Explanation:
The strings s and t can not be made identical as 'o' needs to be mapped to both 'a' and 'r'.

Example 3:
Input: s = "paper", t = "title"
Output: true


Constraints:
1 <= s.length <= 5 * 104
t.length == s.length
s and t consist of any valid ascii character.
*********************************************/

func Output205() any {
	return wordPattern("aaaa", "dog cat cat dog")
}

// * Solution 1 -- Hash Map + Set -- Time: O(n), Space: O(n)
func isIsomorphic(s string, t string) bool {
	sMap := map[rune]rune{}
	added := map[rune]bool{}

	for i, r := range s {
		tmp := rune(t[i])

		if val, found := sMap[r]; !found {
			if added[tmp] {
				return false
			}
			sMap[r] = tmp
			added[tmp] = true
		} else {
			if val != tmp {
				return false
			}
		}
	}
	return true
}

// * Solution 2 -- Two Maps with Byte -- Time: O(n), Space: O(n)
func isIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := map[byte]byte{}
	tMap := map[byte]byte{}

	for i := 0; i < len(s); i++ {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		}
		sMap[s[i]] = byte(i + 1)
		tMap[t[i]] = byte(i + 1)
	}

	return true
}


// * Solution 3 -- Single Map with Index -- Time: O(n), Space: O(n)
func isIsomorphic3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	indexMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		sKey := s[i]
		tKey := t[i] + 128

		if indexMap[sKey] != indexMap[tKey] {
			return false
		}

		indexMap[sKey] = i + 1
		indexMap[tKey] = i + 1
	}

	return true
}

// * Solution 4 -- Array Based -- Time: O(n), Space: O(1)
func isIsomorphic4(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sArr := [256]int{}
	tArr := [256]int{}

	for i := 0; i < len(s); i++ {
		if sArr[s[i]] != tArr[t[i]] {
			return false
		}

		sArr[s[i]] = i + 1
		tArr[t[i]] = i + 1
	}

	return true
}

