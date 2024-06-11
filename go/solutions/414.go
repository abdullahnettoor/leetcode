package solutions

import (
	"math"
)

/* Question *******************************
392. Is Subsequence
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:
Input: s = "abc", t = "ahbgdc"
Output: true

Example 2:
Input: s = "axc", t = "ahbgdc"
Output: false

Constraints:
0 <= s.length <= 100
0 <= t.length <= 104
s and t consist only of lowercase English letters.
******************************************/

func Output414() any {
	return thirdMax([]int{3, 7, 5, 2, 9, 7, 6, 7})
}

// * Solution -- Linear Scan -- Time O(n) - Space O(1)
func thirdMax(nums []int) int {
	m1, m2, m3 := math.MinInt64, math.MinInt64, math.MinInt64
	for i := range nums {
		switch nums[i] {
		case m1, m2, m3:
			continue
		}
		switch {
		case nums[i] > m1:
			m3, m2, m1 = m2, m1, nums[i]
		case nums[i] > m2:
			m3, m2 = m2, nums[i]
		case nums[i] > m3:
			m3 = nums[i]
		}
	}

	if m3 == math.MinInt64 {
		return m1
	}
	return m3
}

// * Solution -- Sorting -- Time O(nlogn) - Space O(logn)
// func thirdMax(nums []int) int {
// 	sort.Ints(nums)
// 	res := []int{nums[len(nums)-1]}
// 	for i := len(nums) - 2; i >= 0; i-- {
// 		if nums[i] == res[len(res)-1] {
// 			continue
// 		}
// 		res = append(res, nums[i])
// 	}
// 	switch len(res) {
// 	case 1, 2:
// 		return res[0]
// 	default:
// 		return res[2]
// 	}
// }
