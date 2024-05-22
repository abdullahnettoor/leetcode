package solutions

/* Question **********************************
? 104. Maximum Depth of Binary Tree
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:
Input: root = [1,null,2]
Output: 2

Constraints:
The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
*********************************************/

func Output104() any {
	l := &TreeNode{Val: 1}
	r := &TreeNode{Val: 50}
	root := &TreeNode{Val: 5, Left: l, Right: r}
	return maxDepth(root)
}

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// * Solution 1 -- Recursion -- DFS -- Time & Space O(n)
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	left := maxDepth(root.Left) + 1
// 	right := maxDepth(root.Right) + 1
// 
// 	if left > right {
// 		return left
// 	}
// 	return right
// }

// * Solution 2 -- Nested Loop -- BFS -- Time & Space O(n)
// func maxDepth(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     queue := []*TreeNode{root}
//     depth := 0
//     for len(queue) != 0 {
//         n := len(queue)
//         for n != 0 {
//             tmp := queue[0]
//             if tmp.Left != nil {
//                 queue = append(queue, tmp.Left)
//             } 
//             if tmp.Right != nil {
//                 queue = append(queue, tmp.Right)
//             } 
//             queue = queue[1:]
//             n--
//         }
//         depth++
//     }
//     return depth
// }

