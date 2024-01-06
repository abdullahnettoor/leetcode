package solutions

/* Question **********************************
? 94. Binary Tree Inorder Traversal
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [1,3,2]

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

func Output94() any {
	l := &TreeNode{Val: 1}
	r := &TreeNode{Val: 50}
	root := &TreeNode{Val: 5, Left: l, Right: r}
	return inorderTraversal(root)
}

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// * Solution 1 Using Recursion
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, inorderTraversal(root.Left)...)
		res = append(res, root.Val)
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}

// * TODO: Solution 2 Using Loop
// func inorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	current := root
// 	stack := []*TreeNode{}
// 	for {
// 		stack = append(stack, current)
// 		if current.Left == nil {
// 			res = append(res, current.Val)
// 			stack = stack[:len(stack)-1]
// 			current = stack[len(stack)-1]
// 		} else {
// 			current = current.Left
// 		}
// 		// res = append(res, current.Val)
// 		if len(stack) == 1 && current != root {
// 			break
// 		}
// 	}
// 	return res
// }
