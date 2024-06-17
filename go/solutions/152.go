package solutions

/* Question **********************************
? 152. Maximum Product Subarray
Given an integer array nums, find a subarray that has the largest product, and return the product.
The test cases are generated so that the answer will fit in a 32-bit integer.

Example 1:
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Example 2:
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 

Constraints:
1 <= nums.length <= 2 * 104
-10 <= nums[i] <= 10
The product of any subarray of nums is guaranteed to fit in a 32-bit integer.
*********************************************/

func Output152() any {
	return maxProductInSubarray([]int{1,2,-3,4-20,9,4,0})
}

// * Solution -- Kadane's Algorithm -- Time O(n) - Space O(1)
func maxProductInSubarray(nums []int) int {
	product := nums[0]
    currMin, currMax := 1, 1
    for _, num := range nums {
		temp := currMax*num
        currMin = min(temp, currMin*num, num)
        currMax = max(temp, currMin*num, num)
        product = max(product, currMax)
    }
    return product
}

// * Solution -- Brute Force - 2 Loops -- Time O(n^2) - Space O(1)
// func maxProductInSubarray(nums []int) int {
//     product := nums[0]
//     for i := range nums {
//         currProduct := 1
//         for j := i; j < len(nums); j++ {
//             currProduct *= nums[j]
//             product = max(product, currProduct)
//         }
//     }
//     return product
// }