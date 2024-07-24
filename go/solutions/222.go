package solutions

/* Question ********************************
? 222. Count Complete Tree Nodes
Given the root of a complete binary tree, return the number of the nodes in the tree.
According to Wikipedia, every level, except possibly the last, is completely filled in a complete binary tree, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

Design an algorithm that runs in less than O(n) time complexity.

Example 1:
Input: root = [1,2,3,4,5,6]
Output: 6

Example 2:
Input: root = []
Output: 0

Example 3:
Input: root = [1]
Output: 1

Constraints:
The number of nodes in the tree is in the range [0, 5 * 104].
0 <= Node.val <= 5 * 104
The tree is guaranteed to be complete.
*****************************************/

func Output222() any {
	return countNodes(nil)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// * Solution 1 -- Recursion -- Time O(n) - Space O(n)
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// * Solution 2 -- Binary Search -- Time O(log^2 n) - Space O(1)
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getLeftHeight(root)
	rightHeight := getRightHeight(root)

	if leftHeight == rightHeight {
		return (1 << leftHeight) - 1
	}

	return 1 + countNodes2(root.Left) + countNodes2(root.Right)
}

func getLeftHeight(node *TreeNode) int {
	height := 0
	for node != nil {
		height++
		node = node.Left
	}
	return height
}

func getRightHeight(node *TreeNode) int {
	height := 0
	for node != nil {
		height++
		node = node.Right
	}
	return height
}

// * Solution 3 -- Level Order Traversal -- Time O(n) - Space O(n)
func countNodes3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		count++

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return count
}
