package solutions

/* Question **********************************
? 230. Kth Smallest Element in a BST
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

Example 1:
Input: root = [3,1,4,null,2], k = 1
Output: 1

Example 2:
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3

Constraints:
The number of nodes in the tree is n.
1 <= k <= n <= 104
0 <= Node.val <= 104

Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?
*********************************************/

func Output230() any {
	l := &TreeNode{Val: 1}
	r := &TreeNode{Val: 50}
	root := &TreeNode{Val: 5, Left: l, Right: r}
	return kthSmallest(root, 3)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// * Solution -- DFS - Recursion -- Time O(H+k) - Space O(1)
func kthSmallest(root *TreeNode, k int) int {
	var res, ctr int
	var inOrder func(node *TreeNode)

	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		ctr++
		if ctr == k {
			res = node.Val
			return
		}
		inOrder(node.Right)
		return
	}

	inOrder(root)

	return res
}

// * Solution -- DFS - Recursion -- Time O(n) - Space O(n)
// func kthSmallest(root *TreeNode, k int) int {
// 	var inOrder func(node *TreeNode, arr []int) []int
// 	inOrder = func(node *TreeNode, arr []int) []int {
// 		if node == nil {
// 			return arr
// 		}
// 		arr = inOrder(node.Left, arr)
// 		arr = append(arr, node.Val)
// 		arr = inOrder(node.Right, arr)
// 		return arr
// 	}
// 	arr := inOrder(root, []int{})
// 	return arr[k-1]
// }

// * Solution -- DFS - Iteration -- Time O(H) - Space O(H)
// * Worst Case will be O(n) for Time & Space for skewed trees
// func kthSmallest(root *TreeNode, k int) int {
// 	stack := []*TreeNode{}
// 	curr := root
// 	for curr != nil || len(stack) > 0 {
// 		if curr != nil {
// 			stack = append(stack, curr)
// 			curr = curr.Left
// 		} else {
// 			node := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			k--
// 			if k == 0 {
// 				curr = node
// 				break
// 			}
// 			curr = node.Right
// 		}
// 	}
// 	return curr.Val
// }
