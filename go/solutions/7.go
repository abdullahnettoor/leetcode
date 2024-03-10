package solutions

import "math"

/* Question ********************************
7. Reverse Integer
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321

Example 3:
Input: x = 120
Output: 21
 
Constraints:
-231 <= x <= 231 - 1
*****************************************/

func Output7() any {
	return reverse(12132432)
}

// * Solution 1 
func reverse(x int) int {
	var rev int
	for x != 0 {
		temp := x % 10
		rev = (rev * 10) + temp
		x = x / 10
		if rev > math.MaxInt32 || rev < math.MinInt32 {
			return 0
		}
	}
	return rev
}
