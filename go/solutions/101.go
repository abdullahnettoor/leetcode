package solutions

/* Question **********************************
? 101. Symmetric Tree
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

Example 1:
Input: root = [1,2,2,3,4,4,3]
Output: true

Example 2:
Input: root = [1,2,2,null,3,null,3]
Output: false

Constraints:
The number of nodes in the tree is in the range [1, 1000].
-100 <= Node.val <= 100

Follow up: Could you solve it both recursively and iteratively?
*********************************************/

func Output101() any {
	l := &TreeNode{Val: 1}
	r := &TreeNode{Val: 50}
	root := &TreeNode{Val: 5, Left: l, Right: r}
	return isSymmetric(root)
}

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// * Solution 1 -- Recursion -- Time & Space O(n)
func isSymmetric(root *TreeNode) bool {
    return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(n1, n2 *TreeNode) bool {
    if n1 == nil && n2 == nil {
        return true
    }
    if n1 == nil || n2 == nil {
        return false
    }
    return n1.Val == n2.Val && isSymmetricHelper(n1.Left, n2.Right) && isSymmetricHelper(n2.Left, n1.Right)
}

