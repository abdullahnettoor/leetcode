package solutions

/* Question ********************************
34. Find First and Last Position of Element in Sorted Array
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Example 3:
Input: nums = [], target = 0
Output: [-1,-1]

Constraints:
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109
*****************************************/

func Output34() any {
	return searchRange([]int{0, 3, 5, 6, 7, 8, 8, 9, 9}, 8)
}

// * Solution 1
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	f, l := 0, len(nums)-1
	for f <= l {
		mid := (f + l) / 2
		if nums[mid] == target {
			f, l = mid, mid
			for f >= 0 && nums[f] == target {
				f--
			}
			for l < len(nums) && nums[l] == target {
				l++
			}
			return []int{f + 1, l - 1}
		} else if nums[mid] < target {
			f = mid + 1
		} else {
			l = mid - 1
		}
	}

	return []int{-1, -1}
}
