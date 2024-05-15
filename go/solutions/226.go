package solutions

/* Question **********************************
? 226. Invert Binary Tree
Given the root of a binary tree, invert the tree, and return its root.

Example 1:
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Example 2:
Input: root = [2,1,3]
Output: [2,3,1]

Example 3:
Input: root = []
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
*********************************************/

func Output226() any {
	l := &TreeNode{Val: 1}
	r := &TreeNode{Val: 50}
	root := &TreeNode{Val: 5, Left: l, Right: r}
	return invertTree(root)
}

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// * Solution 1 -- Recursion -- DFS -- Space & Time O(n)
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}

// * Solution 2 -- Recursion -- DFS -- Space & Time O(n)
// func invertTree(root *TreeNode) *TreeNode {
//     if root == nil {
// 		return root
// 	}
// 	queue := make([]*TreeNode, 0)
// 	queue = append(queue, root)
// 	for len(queue) != 0 {
// 		peekNode := queue[0]
// 		if peekNode.Left != nil {
// 			queue = append(queue, peekNode.Left)
// 		}
// 		if peekNode.Right != nil {
// 			queue = append(queue, peekNode.Right)
// 		}
// 		peekNode.Left, peekNode.Right = peekNode.Right, peekNode.Left
//         queue = queue[1:]
// 	}
// 	return root
// }