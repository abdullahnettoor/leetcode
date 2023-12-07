package main

import (
	"fmt"

	"github.com/abdullahnettoor/leetcode/go/solutions"
)

func main() {
	str := solutions.LongestCommonPrefix([]string{"flower", "flower", "flight"})
	fmt.Println("Solution is", str)
}
