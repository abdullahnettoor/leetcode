package solutions

/* Question **********************************
? 896. Monotonic Array
An array is monotonic if it is either monotone increasing or monotone decreasing.
An array nums is monotone increasing if for all i <= j, nums[i] <= nums[j]. An array nums is monotone decreasing if for all i <= j, nums[i] >= nums[j].

Given an integer array nums, return true if the given array is monotonic, or false otherwise.

Example 1:
Input: nums = [1,2,2,3]
Output: true

Example 2:
Input: nums = [6,5,4,4]
Output: true

Example 3:
Input: nums = [1,3,2]
Output: false

Constraints:
1 <= nums.length <= 10^5
-10^5 <= nums[i] <= 10^5
*********************************************/

func Output896() any {
	return isMonotonic([]int{1, 3, 2})
}

// * Solution 1
func isMonotonic(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	order := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if order == 0 {
				order = 1
			} else if order == -1 {
				return false
			}
		} else if nums[i] < nums[i-1] {
			if order == 0 {
				order = -1
			} else if order == 1 {
				return false
			}
		}
	}
	return true
}
