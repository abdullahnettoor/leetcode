# LeetCode Solutions

A collection of solutions to LeetCode problems, primarily focusing on Go and SQL implementations. More language solutions are planned for the future.

## Project Structure

```markdown
├── README.md
├── go
│   ├── cmd
│   ├── go.mod
│   ├── sample workouts
│   └── solutions
│       └── # Contains all Go solutions
└── sql
```

## Solutions Overview

The repository contains solutions to various LeetCode problems, including:

### Array & String Problems
- Two Sum (#1)
- Longest Common Prefix (#14)
- Valid Parentheses (#20)
- Group Anagrams (#49)
- Matrix Diagonal Sum (#1572)
- Transpose Matrix (#867)
- and more...
### Binary Trees
- Symmetric Tree (#101)
- Maximum Depth of Binary Tree (#104)
- Binary Tree Inorder Traversal (#94)
- Invert Binary Tree (#226)
- Kth Smallest Element in a BST (#230)
- and more...

### Linked Lists
- Merge Two Sorted Lists (#21)
- Reverse Linked List (#206)
- and more...

### Dynamic Programming
- Fibonacci Number (#509)
- Range Sum Query (#303)
- and more...

### Data Structures
- Implement Trie (#208)
- Valid Anagram (#242)
- and more...

## Implementation Details

Each solution includes:
- Problem description with constraints
- Multiple solution approaches where applicable
- Time and space complexity analysis
- Example test cases

## Code Example

Here's a sample solution for the Two Sum problem:

```go
func twoSum(nums []int, target int) []int {
    if len(nums) == 2 {
        return []int{0,1}
    }
    h := make(map[int]int, len(nums))
    for i, num := range nums {
        t := target - num
        if _, ok := h[t]; ok {
            return []int{h[t], i}
        }
        h[num] = i
    }
    return []int{}
}
```

## Running the Solutions

1. Clone the repository:
```bash
git clone https://github.com/abdullahnettoor/leetcode
```

2. Navigate to the project directory:
```bash
cd leetcode/go
```

3. Run a specific solution:
```bash
go run cmd/main.go
```

## Contributing

Contributions are welcome! Feel free to:
- Add new solutions
- Improve existing solutions
- Add solutions in different programming languages
- Improve documentation

## License

This project is open source and available under the MIT License.

## Contact

Created by [@abdullahnettoor](https://github.com/abdullahnettoor)

