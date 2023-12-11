package solutions

import (
	"regexp"
	"strings"
)

/* Question ********************************
125. Valid PalindromeA phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 3:
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:
1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*******************************************/

func IsPalindrome(s string) bool {
	// Remove non-alphanumeric characters and convert to lowercase
	s = strings.ToLower(s)
	// Parse string to remove non alphanumeric characters
	nonAlphaNumericRegx := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = string(nonAlphaNumericRegx.ReplaceAll([]byte(s), []byte("")))

	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
