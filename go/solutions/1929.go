package solutions

import "runtime"

/* Question **********************************
? 1929. Concatenation of Array
Given an integer array nums of length n, you want to create an array ans of length 2n where ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).

Specifically, ans is the concatenation of two nums arrays.
Return the array ans.

Example 1:
Input: nums = [1,2,1]
Output: [1,2,1,1,2,1]
Explanation: The array ans is formed as follows:
- ans = [nums[0],nums[1],nums[2],nums[0],nums[1],nums[2]]
- ans = [1,2,1,1,2,1]

Example 2:
Input: nums = [1,3,2,1]
Output: [1,3,2,1,1,3,2,1]
Explanation: The array ans is formed as follows:
- ans = [nums[0],nums[1],nums[2],nums[3],nums[0],nums[1],nums[2],nums[3]]
- ans = [1,3,2,1,1,3,2,1]

Constraints:
n == nums.length
1 <= n <= 1000
1 <= nums[i] <= 1000
*********************************************/

func Output1929() any {
	return getConcatenation([]int{1, 2, 3, 4, 5, 6})
}

// * Solution 1 -- Slice Appending -- Time: O(n) - Space O(n)
func getConcatenation(nums []int) []int {
	runtime.GC()
	return append(nums, nums...)
}

// * Solution 2 -- Manual Copying -- Time: O(n) - Space O(n)
func getConcatenation2(nums []int) []int {
	n := len(nums)
	result := make([]int, 2*n)

	for i := 0; i < n; i++ {
			result[i], result[i+n] = nums[i], nums[i]
	}

	return result
}