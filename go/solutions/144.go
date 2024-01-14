package solutions

/* Question **********************************
? 144. Binary Tree Preorder Traversal
Given the root of a binary tree, return the preorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [1,2,3]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]

Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

Follow up: Recursive solution is trivial, could you do it iteratively?
*********************************************/

func Output144() any {
	l := &TreeNode{Val: 5}
	r := &TreeNode{Val: 3}
	root := &TreeNode{Val: 1, Left: l, Right: r}
	return preorderTraversal(root)
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
func preorderTraversal(root *TreeNode) []int {
	res := &[]int{}
	if root != nil {
		l := preorderTraversal(root.Left)
		r := preorderTraversal(root.Right)
		*res = append(*res, root.Val)
		*res = append(*res, l...)
		*res = append(*res, r...)
	}
	return *res
}
