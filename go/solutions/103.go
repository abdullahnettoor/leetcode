package solutions

/* Question ********************************
? 103. Binary Tree Zigzag Level Order Traversal
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100
*****************************************/

func Output103() any {
	return zigzagLevelOrder(nil)
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	isLtr := true

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if isLtr {
				level[i] = node.Val
			} else {
				level[levelSize-1-i] = node.Val
			}
		}
		isLtr = !isLtr
		result = append(result, level)
	}

	return result
}

// * Solution 2 -- DFS (Preorder) -- Time O(n) - Space O(h)
func zigzagLevelOrder2(root *TreeNode) [][]int {
	result := [][]int{}

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level >= len(result) {
			result = append(result, []int{})
		}

		if level%2 == 0 {
			result[level] = append(result[level], node.Val)
		} else {
			result[level] = append([]int{node.Val}, result[level]...)
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)
	return result
}

