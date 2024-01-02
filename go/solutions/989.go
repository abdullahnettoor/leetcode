package solutions

import (
	"fmt"
	"math/big"
	"strconv"
)

/* Question **********************************
? 989. Add to Array-Form of Integer
The array-form of an integer num is an array representing its digits in left to right order.

For example, for num = 1321, the array form is [1,3,2,1].
Given num, the array-form of an integer, and an integer k, return the array-form of the integer num + k.

Example 1:
Input: num = [1,2,0,0], k = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234

Example 2:
Input: num = [2,7,4], k = 181
Output: [4,5,5]
Explanation: 274 + 181 = 455

Example 3:
Input: num = [2,1,5], k = 806
Output: [1,0,2,1]
Explanation: 215 + 806 = 1021

Constraints:
1 <= num.length <= 104
0 <= num[i] <= 9
num does not contain any leading zeros except for the zero itself.
1 <= k <= 104
*********************************************/

func Output989() any {
	return addToArrayForm([]int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}, 806)
}

// * Solution 1
func addToArrayForm(num []int, k int) []int {
	var s string
	for i := range num {
		s += strconv.Itoa(num[i])
	}
	n, _ := new(big.Int).SetString(s, 10)
	n.Add(n, new(big.Int).SetInt64(int64(k)))
	s = fmt.Sprint(n)
	res := make([]int, len(s))
	for i := range res {
		res[i], _ = strconv.Atoi(string(s[i]))
	}
	return res
}

// TODO: Solution without using bigInt
