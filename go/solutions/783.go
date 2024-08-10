package solutions

import (
	"math"
	"sort"
)

/* Question ********************************
? 783. Minimum Distance Between BST Nodes
Given the root of a Binary Search Tree (BST), return the minimum difference between the values of any two different nodes in the tree.

Example 1:
Input: root = [4,2,6,1,3]
Output: 1

Example 2:
Input: root = [1,0,48,null,null,12,49]
Output: 1

Constraints:
The number of nodes in the tree is in the range [2, 100].
0 <= Node.val <= 105

Note: This question is the same as 530.go: https://leetcode.com/problems/minimum-absolute-difference-in-bst/
******************************************/

func Output783() any {
	return minDiffInBST(nil)
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
func minDiffInBST(root *TreeNode) int {

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

// * Solution 2 -- DFS (Inorder) with prev pointer -- Time O(n) - Space O(h)
func minDiffInBST2(root *TreeNode) int {
	minimum := math.MaxInt32
	var prev *TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		if prev != nil {
			minimum = min(minimum, node.Val-prev.Val)
		}
		prev = node

		dfs(node.Right)
	}

	dfs(root)
	return minimum
}

// * Solution 3 -- BFS (Level Order) -- Time O(n) - Space O(w)
func minDiffInBST3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minimum := math.MaxInt32
	values := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		values = append(values, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	sort.Ints(values)
	for i := 1; i < len(values); i++ {
		minimum = min(minimum, values[i]-values[i-1])
	}

	return minimum
}

// * Solution 4 -- Morris Traversal -- Time O(n) - Space O(1)
func minDiffInBST4(root *TreeNode) int {
	minimum := math.MaxInt32
	var prev *TreeNode
	current := root

	for current != nil {
		if current.Left == nil {
			if prev != nil {
				minimum = min(minimum, current.Val-prev.Val)
			}
			prev = current
			current = current.Right
		} else {
			// Find the inorder predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				predecessor.Right = current
				current = current.Left
			} else {
				predecessor.Right = nil
				if prev != nil {
					minimum = min(minimum, current.Val-prev.Val)
				}
				prev = current
				current = current.Right
			}
		}
	}

	return minimum
}
