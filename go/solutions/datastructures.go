package solutions

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}
