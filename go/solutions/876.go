package solutions

/* Question **********************************
? 876. Middle of the Linked List
Given the head of a singly linked list, return the middle node of the linked list.
If there are two middle nodes, return the second middle node.

Example 1:
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.

Constraints:
The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100
*********************************************/

func Output876() any {
	l2 := &ListNode{Val: 1}
	l1 := &ListNode{Val: 50, Next: l2}
	return middleNode(l1)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// * Solution -- Iterative -- Time O(n) - Space O(1)
func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

// * Solution -- Recursion -- Time & Space O(n/2)
// func middleNode(head *ListNode) *ListNode {
//     return middleNodeHelper(head.Next, head)
// }

// func middleNodeHelper(fast, slow *ListNode) *ListNode {
//     if fast == nil {d
//         return slow
//     } 
//     if fast.Next == nil {
//         return slow.Next
//     }
//     return middleNodeHelper(fast.Next.Next, slow.Next)
// }

