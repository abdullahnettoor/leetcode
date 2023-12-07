package leetcode

// 1796. Second Largest Digit in a String

// Given an alphanumeric string s, return the second largest numerical digit that appears in s, or -1 if it does not exist.
// An alphanumeric string is a string consisting of lowercase English letters and digits.

// Example 1:
// Input: s = "dfa12321afd"
// Output: 2
// Explanation: The digits that appear in s are [1, 2, 3]. The second largest digit is 2.

// Example 2:
// Input: s = "abc1111"
// Output: -1
// Explanation: The digits that appear in s are [1]. There is no second largest digit.

// Constraints:
// 1 <= s.length <= 500
// s consists of only lowercase English letters and/or digits.

func secondHighest(s string) int {
	if len(s) == 1 {
		return -1
	}

	var largest, secondLargest rune

	for _, value := range s {
		if value >= '0' && value <= '9' {
			if value > largest {
				secondLargest = largest
				largest = value
			} else if value >= secondLargest && value != largest {
				secondLargest = value
			}
		}
	}

	if secondLargest >= '0' && secondLargest <= '9' {
		return int(secondLargest) - 48
	} else {
		return -1
	}

}
