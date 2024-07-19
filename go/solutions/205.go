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

