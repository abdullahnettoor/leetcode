package solutions

import "strings"

/* Question ********************************
? 151. Reverse Words in a String
Given an input string s, reverse the order of the words.
A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

Example 1:
Input: s = "the sky is blue"
Output: "blue is sky the"

Example 2:
Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:
Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.


Constraints:
1 <= s.length <= 104
s contains English letters (upper-case and lower-case), digits, and spaces ' '.
There is at least one word in s.

Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?
*****************************************/

func Output151() any {
	return reverseWords("hi there")
}

// * Solution 1 -- Split and Reverse -- Time O(n) - Space O(n)
func reverseWords(s string) string {
	words := []string{}

	word := ""
	for _, r := range s {
		if r == ' ' {
			if word != "" {
				words = append([]string{word}, words...)
			}
			word = ""
		} else {
			word += string(r)
		}
	}
	if word != "" {
		words = append([]string{word}, words...)
	}

	return strings.Join(words, " ")
}


// * Solution 2 -- Two Pointers -- Time O(n) - Space O(n)
func reverseWords2(s string) string {
	bytes := []byte(strings.TrimSpace(s))
	n := len(bytes)

	reverse := func(start, end int) {
		for start < end {
			bytes[start], bytes[end] = bytes[end], bytes[start]
			start++
			end--
		}
	}
	reverse(0, n-1)

	start := 0
	for i := 0; i < n; i++ {
		if bytes[i] == ' ' {
			reverse(start, i-1)
			start = i + 1
		} else if i == n-1 {
			reverse(start, i)
		}
	}

	j := 0
	for i := 0; i < n; i++ {
		if !(bytes[i] == ' ' && (i == 0 || bytes[i-1] == ' ')) {
			bytes[j] = bytes[i]
			j++
		}
	}

	return string(bytes[:j])
}
