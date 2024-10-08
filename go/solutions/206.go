package solutions

/* Question ********************************
? 206. Reverse Linked List
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []

Constraints:
The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000

Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
*******************************************/

func Output206() any {
	return titleToNumber("AA")
}

/**
Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
*/

// * Solution -- Iterative -- Time O(n) = Space O(1)
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	var prev *ListNode
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
