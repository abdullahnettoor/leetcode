package solutions

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

type Node struct {
	Val       int
	Neighbors []*Node
}
