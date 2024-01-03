package solutions

/* Question **********************************
? 344. Reverse String
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

Example 1:
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]

Constraints:
1 <= s.length <= 105
s[i] is a printable ascii character.
*********************************************/

func Output344() any {
	return reverseString([]string{"H", "a", "n", "n", "a", "h"})
}

// * Solution 1
func reverseString(s []string) []string {
	l, r := 0, len(s)-1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return s
}
