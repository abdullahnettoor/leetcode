package solutions

/* Question **********************************
? 700. Search in a Binary Search Tree
You are given the root of a binary search tree (BST) and an integer val.

Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.

Example 1:
Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]

Example 2:
Input: root = [4,2,7,1,3], val = 5
Output: []


Constraints:
The number of nodes in the tree is in the range [1, 5000].
1 <= Node.val <= 107
root is a binary search tree.
1 <= val <= 107
*********************************************/

func Output700() any {
	return searchBST(nil, 0)
}

// * Solution -- Iterative -- Time O(logn) - Space O(1)
func searchBST(root *TreeNode, val int) *TreeNode {
	curr := root

	for curr != nil {
		switch {
		case curr.Val == val:
			return curr
		case curr.Val > val:
			curr = curr.Left
		default:
			curr = curr.Right
		}
	}

	return nil
}
