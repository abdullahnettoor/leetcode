package solutions

/* Question ********************************
? 199. Binary Tree Right Side View
Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Example 2:
Input: root = [1,null,3]
Output: [1,3]

Example 3:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
*****************************************/

func Output199() any {
	return rightSideView(nil)
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	result := []int{}

	for len(queue) > 0 {
		levelSize := len(queue)
		var rightMostNode *TreeNode

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			rightMostNode = node
		}

		result = append(result, rightMostNode.Val)
	}
	return result
}

// * Solution 2 -- DFS (Preorder) -- Time O(n) - Space O(h)
func rightSideView2(root *TreeNode) []int {
	result := []int{}

	var dfs func(node *TreeNode, depth int, result *[]int)
	dfs = func(node *TreeNode, depth int, result *[]int) {
		if node == nil {
			return
		}

		if depth == len(*result) {
			*result = append(*result, node.Val)
		}

		dfs(node.Right, depth+1, result)
		dfs(node.Left, depth+1, result)
	}

	dfs(root, 0, &result)
	return result
}
