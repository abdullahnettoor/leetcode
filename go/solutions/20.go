package solutions

/* Question **********************************
? 20. Valid Parentheses
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Constraints:
1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*********************************************/

func Output20() any {
	return isValid("[])")
}

// * Solution 1 -- Using Stack
func isValid(s string) bool {
	chars := []rune{}
	for i := range s {
		switch rune(s[i]) {
		case '(', '{', '[':
			chars = append(chars, rune(s[i]))
		case ')', '}', ']':
			if len(chars) != 0 {
				switch rune(s[i]) {
				case ')':
					if chars[len(chars)-1] != '(' {
						return false
					}
				case '}':
					if chars[len(chars)-1] != '{' {
						return false
					}
				case ']':
					if chars[len(chars)-1] != '[' {
						return false
					}
				}
				chars = chars[:len(chars)-1]
			} else {
				return false
			}
		}
	}
	return len(chars) == 0
}

// * Solution 2 -- Old Solution
// func isValid(s string) bool {
// 	length := len(s)
// 	if length < 2 {
// 		return false
// 	}
// 	if s[0] == ']' || s[0] == '}' || s[0] == ')' {
// 		return false
// 	}
// 	if s[length-1] == '[' || s[length-1] == '{' || s[length-1] == '(' {
// 		return false
// 	}

// 	chars := []rune{}
// 	for i := range s {

// 		switch rune(s[i]) {
// 		case '(', '{', '[':
// 			chars = append(chars, rune(s[i]))

// 		case ')', '}', ']':
// 			if len(chars) != 0 {
// 				switch rune(s[i]) {
// 				case ')':
// 					if chars[len(chars)-1] == '(' {
// 						chars = chars[:len(chars)-1]
// 					} else {
// 						return false
// 					}
// 				case '}':
// 					if chars[len(chars)-1] == '{' {
// 						chars = chars[:len(chars)-1]
// 					} else {
// 						return false
// 					}
// 				case ']':
// 					if chars[len(chars)-1] == '[' {
// 						chars = chars[:len(chars)-1]
// 					} else {
// 						return false
// 					}
// 				}
// 			} else {
// 				return false
// 			}
// 		}
// 	}
// 	return len(chars) == 0
// }
