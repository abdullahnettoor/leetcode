package solutions

/* Question ********************************
? 53. Maximum Subarray
Given an integer array nums, find the subarray with the largest sum, and return its sum.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:
Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

Constraints:
1 <= nums.length <= 105
-104 <= nums[i] <= 104

Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*****************************************/

func Output53() any {
	return maxSubArray([]int{1, 1, -1, 3, 3, 4, 3, -2, -4, 2})
}

// * Solution -- Kadane's Algorithm -- Time O(n) - Space O(1)
func maxSubArray(nums []int) int {
	maxSum, currSum := nums[0], 0
	for _, num := range nums {
		currSum += num
		maxSum = max(maxSum, currSum)
		if currSum < 0 {
			currSum = 0
		}
	}
	return maxSum
}

// * Solution -- Brute Force - 2 Loops -- Time O(n^2) - Space O(1)
// func maxSubArray(nums []int) int {
//     maxSum := nums[0]
//     for i := range nums {
//         sum := 0
//         for j := i; j < len(nums); j++ {
//             sum += nums[k]
//             maxSum = max(sum, maxSum)
//       }
//     }
//     return maxSum
// }

// * Solution -- Brute Force - 3 Loops -- Time O(n^3) - Space O(1)
// func maxSubArray(nums []int) int {
//     maxSum := nums[0]
//     for i := range nums {
//       for j := i; j < len(nums); j++ {
//         sum := 0
//         for k := i; k <= j; k++ {
//             sum += nums[k]
//             maxSum = max(sum, maxSum)
//         }
//       }
//     }
//     return maxSum
// }
