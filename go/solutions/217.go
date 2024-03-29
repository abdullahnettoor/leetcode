package solutions

/* Question ********************************
217. Contains Duplicate
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:
Input: nums = [1,2,3,1]
Output: true

Example 2:
Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

Constraints:
1 <= nums.length <= 105
-109 <= nums[i] <= 109
*****************************************/

func Output217() any {
	return containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
}

// * Solution 1 - Using hast table (map)
func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
		if m[n] > 1 {
			return true
		}
	}
	return false
}

// * Solution 2 - Using 2 loops
// func containsDuplicate(nums []int) bool {
// 	for i := range nums {
// 		length := len(nums) - 1
// 		for length != i {
// 			if nums[i] == nums[length] {
// 				return true
// 			}
// 			length -= 1
// 		}
// 	}
// 	return false
// }
