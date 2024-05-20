package solutions

/* Question **********************************
? 108. Convert Sorted Array to Binary Search Tree
Given an integer array nums where the elements are sorted in ascending order, convert it to a
height-balanced binary search tree.

Example 1:
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:

Example 2:
Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.


Constraints:
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in a strictly increasing order.
*********************************************/

func Output704() any {
	return search([]int{5, 10, 11, 14, 25, 30}, 10)
}

// * Solution 1 -- Divide & Conquer -- Using Loop -- Time O(logn) - Space O(1)
func search(nums []int, target int) int {
    start, end := 0, len(nums)-1
    for start <= end {
        mid := start + (end-start)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            start = mid+1
        } else {
            end = mid-1
        }
    }
    return -1
}


// * Solution 2 -- Recursion -- Using helper function -- Time O(logn) -- Space O(n)
// func search(nums []int, target int) int {
//     return searchHelper(0, len(nums)-1, target, nums)
// }

// func searchHelper(start, end, target int, nums []int) int {
//     if start > end {
//         return -1
//     }
//     mid := (start + end)/ 2
//     if nums[mid] == target {
//         return mid
//     } else if nums[mid] > target {
//         return searchHelper(start, mid-1, target, nums)
//     } else {
//         return searchHelper(mid+1, end, target, nums)
//     }
// }
