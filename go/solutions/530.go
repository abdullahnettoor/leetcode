package solutions

import (
	"math"
)

/* Question ********************************
? 530. Minimum Absolute Difference in BST
Given the root of a Binary Search Tree (BST), return the minimum absolute difference between the values of any two different nodes in the tree.

Example 1:
Input: root = [4,2,6,1,3]
Output: 1

Example 2:
Input: root = [1,0,48,null,null,12,49]
Output: 1

Constraints:
The number of nodes in the tree is in the range [2, 104].
0 <= Node.val <= 105

Note: This question is the same as 783: https://leetcode.com/problems/minimum-distance-between-bst-nodes/
******************************************/

func Output530() any {
	return getMinimumDifference(nil)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// * Solution 1 -- DFS (Inorder) -- Time O(n) - Space O(n)
func getMinimumDifference(root *TreeNode) int {

	arr := []int{}

	var dfs func(node *TreeNode, a *[]int)
	dfs = func(node *TreeNode, a *[]int) {
		if node == nil {
			return
		}
		if node.Left != nil {
			dfs(node.Left, a)
		}
		*a = append(*a, node.Val)
		if node.Right != nil {
			dfs(node.Right, a)
		}
	}

	dfs(root, &arr)

	minimum := math.MaxInt32
	for i := 1; i < len(arr); i++ {
		minimum = min(minimum, (arr[i] - arr[i-1]))
	}

	return minimum
}
