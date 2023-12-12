package solutions

/* Question **********************************
1464. Maximum Product of Two Elements in an Array
Easy
Topics
Companies
Hint
Given the array of integers nums, you will choose two different indices i and j of that array. Return the maximum value of (nums[i]-1)*(nums[j]-1).


Example 1:
Input: nums = [3,4,5,2]
Output: 12
Explanation: If you choose the indices i=1 and j=2 (indexed from 0), you will get the maximum value, that is, (nums[1]-1)*(nums[2]-1) = (4-1)*(5-1) = 3*4 = 12.

Example 2:
Input: nums = [1,5,4,5]
Output: 16
Explanation: Choosing the indices i=1 and j=3 (indexed from 0), you will get the maximum value of (5-1)*(5-1) = 16.

Example 3:
Input: nums = [3,7]
Output: 12

Constraints:
2 <= nums.length <= 500
1 <= nums[i] <= 10^3
*********************************************/

func maxProduct(nums []int) int {
	var l1, l2 int
	for i := range nums {
		if nums[i] >= l1 {
			l2 = l1
			l1 = nums[i]
		} else if nums[i] >= l2 {
			l2 = nums[i]
		}
	}
	return (l2 - 1) * (l1 - 1)
}

// func MaxProduct(nums []int) int {
// 	size := len(nums)
// 	sort.Ints(nums)
// 	return (nums[size-1] - 1) * (nums[size-2] - 1)
// }

func Output1464() any {
	return maxProduct([]int{3, 4, 5, 5, 4, 2})
}
