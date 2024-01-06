package solutions

/* Question **********************************
? 450. Delete Node in a BST
Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:
Search for a node to remove.
If the node is found, delete the node.


Example 1:
Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]
Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.

Example 2:
Input: root = [5,3,6,2,4,null,7], key = 0
Output: [5,3,6,2,4,null,7]
Explanation: The tree does not contain a node with value = 0.

Example 3:
Input: root = [], key = 0
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 104].
-10^5 <= Node.val <= 10^5
Each node has a unique value.
root is a valid binary search tree.
-105 <= key <= 105
*********************************************/

func Output450() any {
	l := &TreeNode{Val: 1}
	r := &TreeNode{Val: 50}
	root := &TreeNode{Val: 5, Left: l, Right: r}
	return deleteNode(root, 1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// * Solution 1
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Right == nil && root.Left == nil {
			root = nil
		} else if root.Right == nil {
			root = root.Left
		} else if root.Left == nil {
			root = root.Right
		} else {
			successor := minNode(root.Right)
			root.Val = successor.Val
			root.Right = deleteNode(root.Right, successor.Val)
		}
	}
	return root
}

func minNode(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
