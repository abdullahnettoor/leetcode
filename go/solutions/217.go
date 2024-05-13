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

// * Solution 1 -- Using hast table -- Time O(n) - Space O(n)
func containsDuplicate(nums []int) bool {
    m := map[int]struct{}{}
    for _, n := range nums {
        if _, found := m[n]; found {
            return true
        }
        m[n] = struct{}{}
    }
    return false
}

// * Solution 2 -- Using Two Pointers -- Sliding Window -- Time O(nlogn) - Space O(1)
// func containsDuplicate(nums []int) bool {
//     sort.Ints(nums)
//     ptr1, ptr2 := 0, 1
//     for ptr2 < len(nums) {
//         if nums[ptr1] == nums[ptr2] {
//             return true
//         }
//         ptr1++
//         ptr2++
//     }
//     return false
// }

// * Solution 3 -- Using Nested Loop -- Brute Force -- Time O(n^2) - Space O(1)
// func containsDuplicate(nums []int) bool {
// 	for left := range nums {
// 		right := len(nums)-1
// 		for left != right {
// 			if nums[left] == nums[right]{
// 				return true
// 			}
// 			right--
// 		}
// 	}
// 	return false
// }
