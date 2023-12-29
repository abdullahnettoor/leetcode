package solutions

/* Question **********************************
? 507. Perfect Number
A perfect number is a positive integer that is equal to the sum of its positive divisors, excluding the number itself. A divisor of an integer x is an integer that can divide x evenly.

Given an integer n, return true if n is a perfect number, otherwise return false.

Example 1:
Input: num = 28
Output: true
Explanation: 28 = 1 + 2 + 4 + 7 + 14
1, 2, 4, 7, and 14 are all divisors of 28.

Example 2:
Input: num = 7
Output: false

Constraints:
1 <= num <= 108
*********************************************/

func Output507() any {
	return checkPerfectNumber(2343)
}

// * Solution 1
// * There is only following perfect numbers until 10^8
func checkPerfectNumber(num int) bool {
	return num == 6 || num == 28 || num == 496 || num == 8128 || num == 33550336
}

// * Solution 2
// * Right Solution, but need more time
// func checkPerfectNumber(num int) bool {
//     tmp := num
//     for i := num-1; i > 0; i--{
//         if num % i == 0 {
//             tmp -= i
//         }
//     }
//     return tmp == 0
// }
