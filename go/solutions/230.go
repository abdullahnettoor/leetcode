package solutions

/* Question **********************************
? 230. Reverse String
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

Example 1:
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]

Constraints:
1 <= s.length <= 105
s[i] is a printable ascii character.
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

// * Solution -- DFS - Recursion -- Time O(n) - Space O(n)
func kthSmallest(root *TreeNode, k int) int {
	var inOrder func(node *TreeNode, arr []int) []int
	inOrder = func(node *TreeNode, arr []int) []int {
		if node == nil {
			return arr
		}
		arr = inOrder(node.Left, arr)
		arr = append(arr, node.Val)
		arr = inOrder(node.Right, arr)
		return arr
	}

	arr := inOrder(root, []int{})

	return arr[k-1]
}
