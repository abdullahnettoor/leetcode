package solutions

/* Question **********************************
? 965. Univalued Binary Tree
A binary tree is uni-valued if every node in the tree has the same value.
Given the root of a binary tree, return true if the given tree is uni-valued, or false otherwise.

Example 1:
Input: root = [1,1,1,1,1,null,1]
Output: true

Example 2:
Input: root = [2,2,2,5,2]
Output: false

Constraints:
The number of nodes in the tree is in the range [1, 100].
0 <= Node.val < 100
*********************************************/

func Output965() any {
	left := &TreeNode{Val: 1}
	right := &TreeNode{Val: 1}
	root := &TreeNode{Val: 1, Left: left, Right: right}
	return isUnivalTree(root)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// * Solution 1
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Right != nil && root.Val != root.Right.Val {
		return false
	}
	if root.Left != nil && root.Val != root.Left.Val {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
