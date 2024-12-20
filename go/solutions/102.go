package solutions

/* Question ********************************
? 102. Binary Tree Level Order Traversal
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000
*****************************************/

func Output102() any {
	return levelOrder(nil)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// * Solution 1 -- Level Order Traversal -- Time O(n) - Space O(n)
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		level := []int{}

		for range queue {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			level = append(level, node.Val)
		}

		result = append(result, level)
	}
	return result
}

// * Solution 2 -- DFS (Preorder) -- Time O(n) - Space O(h)
func levelOrder2(root *TreeNode) [][]int {
	result := [][]int{}

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level >= len(result) {
			result = append(result, []int{})
		}

		result[level] = append(result[level], node.Val)

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)
	return result
}
