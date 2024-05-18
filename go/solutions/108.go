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

func Output108() any {
	return sortedArrayToBST([]int{5, 10, 11, 14, 25, 30})
}

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// * Solution 1 -- Divide & Conquer -- Recursion Using Single Function -- Time & Space O(n)
func sortedArrayToBST(nums []int) *TreeNode {
    n := len(nums)
    if n < 1 {
        return nil
    }

    mid := n/2

    node := &TreeNode{Val:nums[mid]}
    node.Left = sortedArrayToBST(nums[:mid])
    node.Right = sortedArrayToBST(nums[mid+1:])

    return node
}

// * Solution 2 -- Recursion -- Using helper function -- Time & Space O(n)
// func sortedArrayToBST(nums []int) *TreeNode {
// 	return sortedArrayToBSTHelper(0, len(nums)-1, nums)
// }

// func sortedArrayToBSTHelper(start, end int, nums []int) *TreeNode {
// 	if start > end {
// 		return nil
// 	}
// 	mid := (start + end) / 2
// 	node := &TreeNode{Val: nums[mid]}
// 	node.Left = sortedArrayToBSTHelper(start, mid-1, nums)
// 	node.Right = sortedArrayToBSTHelper(mid+1, end, nums)
// 	return node
// }
