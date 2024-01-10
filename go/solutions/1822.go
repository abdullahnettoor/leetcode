package solutions

/* Question **********************************
? 1822. Sign of the Product of an Array
There is a function signFunc(x) that returns:

1 if x is positive.
-1 if x is negative.
0 if x is equal to 0.
You are given an integer array nums. Let product be the product of all values in the array nums.

Return signFunc(product).

Example 1:
Input: nums = [-1,-2,-3,-4,3,2,1]
Output: 1
Explanation: The product of all values in the array is 144, and signFunc(144) = 1

Example 2:
Input: nums = [1,5,0,2,-3]
Output: 0
Explanation: The product of all values in the array is 0, and signFunc(0) = 0

Example 3:
Input: nums = [-1,1,-1,1,-1]
Output: -1
Explanation: The product of all values in the array is -1, and signFunc(-1) = -1

Constraints:
1 <= nums.length <= 1000
-100 <= nums[i] <= 100
*********************************************/

func Output1822() any {
	return arraySign([]int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24})
}

// * Solution 1
func arraySign(nums []int) int {
	negative := 0
	for i := range nums {
		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			negative++
		}
	}
	if negative%2 == 0 {
		return 1
	}
	return -1
}
