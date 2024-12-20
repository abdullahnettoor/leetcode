package solutions

import "sort"

/* Question **********************************
? 2215. Find the Difference of Two Arrays
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.

Note that the integers in the lists may be returned in any order.

Example 1:
Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Explanation:
For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].

Example 2:
Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
Output: [[3],[]]
Explanation:
For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
Every integer in nums2 is present in nums1. Therefore, answer[1] = [].

Constraints:
1 <= nums1.length, nums2.length <= 1000
-1000 <= nums1[i], nums2[i] <= 1000
*********************************************/

func Output2215() any {
	return findDifference([]int{9, 1, 3, 5, 6}, []int{2, 1, 8, 5, 6})
}

// * Solution 1 -- Hashmap -- Time O(m+n) - Space O(m+n)
func findDifference(nums1 []int, nums2 []int) [][]int {
	num1Map := map[int]bool{}
	num2Map := map[int]bool{}

	diff1 := make([]int, 0)
	diff2 := make([]int, 0)

	for _, n := range nums1 {
		num1Map[n] = true
	}
	for _, n := range nums2 {
		num2Map[n] = true
	}

	for _, n := range nums1 {
		if !num2Map[n] {
			diff1 = append(diff1, n)
			num2Map[n] = true
		}
	}
	for _, n := range nums2 {
		if !num1Map[n] {
			diff2 = append(diff2, n)
			num1Map[n] = true
		}
	}

	return [][]int{diff1, diff2}
}

// * Solution 2 -- Using Set -- Time O(m+n) - Space O(m+n)
func findDifference2(nums1 []int, nums2 []int) [][]int {
	set1 := map[int]bool{}
	set2 := map[int]bool{}

	for _, n := range nums1 {
		set1[n] = true
	}
	for _, n := range nums2 {
		set2[n] = true
	}

	diff1 := []int{}
	diff2 := []int{}

	for n := range set1 {
		if !set2[n] {
			diff1 = append(diff1, n)
		}
	}
	for n := range set2 {
		if !set1[n] {
			diff2 = append(diff2, n)
		}
	}

	return [][]int{diff1, diff2}
}

// * Solution 3 -- Sorting + Two Pointers -- Time O(m+n) - Space O(m+n)
func findDifference3(nums1 []int, nums2 []int) [][]int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0
	diff1, diff2 := []int{}, []int{}

	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && (j == len(nums2) || nums1[i] < nums2[j]) {
			if len(diff1) == 0 || nums1[i] != diff1[len(diff1)-1] {
				diff1 = append(diff1, nums1[i])
			}
			i++
		} else if j < len(nums2) && (i == len(nums1) || nums2[j] < nums1[i]) {
			if len(diff2) == 0 || nums2[j] != diff2[len(diff2)-1] {
				diff2 = append(diff2, nums2[j])
			}
			j++
		} else {
			i++
			j++
		}
	}

	return [][]int{diff1, diff2}
}
