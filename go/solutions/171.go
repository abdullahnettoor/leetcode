package solutions

import (
	"math"
)

/* Question ********************************
? 171. Excel Sheet Column Number
Easy
Topics
Companies
Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.

For example:

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...

Example 1:
Input: columnTitle = "A"
Output: 1

Example 2:
Input: columnTitle = "AB"
Output: 28

Example 3:
Input: columnTitle = "ZY"
Output: 701

Constraints:
1 <= columnTitle.length <= 7
columnTitle consists only of uppercase English letters.
columnTitle is in the range ["A", "FXSHRXW"].
*******************************************/

func Output171() any {
	return titleToNumber("AA")
}

func titleToNumber(s string) int {
	var sum int
	for pos, idx := 0, len(s)-1; idx >= 0; idx-- {
		sum += int(math.Pow(26, float64(pos))) * (int(s[idx]) - 'A' + 1)
		pos++
	}
	return sum
}
