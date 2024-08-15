package solutions

/* Question **********************************
? 2390. Removing Stars From a String
You are given a string s, which contains stars *.

In one operation, you can:
Choose a star in s.
Remove the closest non-star character to its left, as well as remove the star itself.
Return the string after all stars have been removed.

Note:
The input will be generated such that the operation is always possible.
It can be shown that the resulting string will always be unique.

Example 1:
Input: s = "leet**cod*e"
Output: "lecoe"
Explanation: Performing the removals from left to right:
- The closest character to the 1st star is 't' in "leet**cod*e". s becomes "lee*cod*e".
- The closest character to the 2nd star is 'e' in "lee*cod*e". s becomes "lecod*e".
- The closest character to the 3rd star is 'd' in "lecod*e". s becomes "lecoe".
There are no more stars, so we return "lecoe".

Example 2:
Input: s = "erase*****"
Output: ""
Explanation: The entire string is removed, so we return an empty string.

Constraints:
1 <= s.length <= 105
s consists of lowercase English letters and stars *.
The operation above can be performed on s.
*********************************************/

func Output2390() any {
	return removeStars("de*fd**fdfw")
}

// * Solution 1 -- Stack -- Time: O(n) - Space: O(n)
func removeStars(s string) string {
	res := make([]rune, 0, len(s))

	for _, r := range s {
		if r == '*' {
			res = res[:len(res)-1]
		} else {
			res = append(res, r)
		}
	}

	return string(res)
}


// * Solution 2 -- Pointers -- Time: O(n) - Space: O(n)
func removeStars2(s string) string {
	res := make([]rune, len(s))
	pos := 0

	for _, r := range s {
		if r == '*' {
			pos--
		} else {
			res[pos] = r
			pos++
		}
	}

	return string(res[:pos])
}

// * Solution 3 -- Update In Place -- Time: O(n) - Space: O(1)
func removeStars3(s string) string {
	runes := []rune(s)
	pos := 0

	for _, r := range runes {
		if r == '*' {
			pos--
		} else {
			runes[pos] = r
			pos++
		}
	}

	return string(runes[:pos])
}