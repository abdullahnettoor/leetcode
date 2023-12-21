package solutions

/* Question **********************************
? 342. Power of Four
Given an integer n, return true if it is a power of four. Otherwise, return false.
An integer n is a power of four, if there exists an integer x such that n == 4x.

Example 1:
Input: n = 16
Output: true

Example 2:
Input: n = 5
Output: false

Example 3:
Input: n = 1
Output: true

Constraints:
-2^31 <= n <= 2^31 - 1

Follow up: Could you solve it without loops/recursion?
*********************************************/

func Output342() any {
	return isPowerOfFour(6)
}

// * Solution 1
func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}
	if n <= 0 || n%4 != 0 {
		return false
	}
	return isPowerOfFour(n / 4)
}

// * Solution 2
// func isPowerOfFour(n int) bool {
// 	return findPower(1, n)
// }

// func findPower(v, n int) bool {
// 	if v == n {
// 		return true
// 	}
// 	if v > n {
// 		return false
// 	}
// 	return findPower(v*4, n)
// }
