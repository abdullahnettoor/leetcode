package solutions

/* Question **********************************
? 2710. Remove Trailing Zeros From a String
Given a positive integer num represented as a string, return the integer num without trailing zeros as a string.

Example 1:
Input: num = "51230100"
Output: "512301"
Explanation: Integer "51230100" has 2 trailing zeros, we remove them and return integer "512301".

Example 2:
Input: num = "123"
Output: "123"
Explanation: Integer "123" has no trailing zeros, we return integer "123".

Constraints:
1 <= num.length <= 1000
num consists of only digits.
num doesn't have any leading zeros.
*********************************************/

func Output2710() any {
	return removeTrailingZeros("12300100")
}

// * Solution 1
func removeTrailingZeros(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] != '0' {
			return num[:i+1]
		}
	}
	return num
}

// * Solution 2
// func removeTrailingZeros(num string) string {
// 	return strings.TrimRight(num, "0")
// }