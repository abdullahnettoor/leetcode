package solutions

import "strings"

/* Question **********************************
? 557. Reverse Words in a String III
Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:
Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"

Example 2:
Input: s = "Mr Ding"
Output: "rM gniD"
 

Constraints:
1 <= s.length <= 5 * 104
s contains printable ASCII characters.
s does not contain any leading or trailing spaces.
There is at least one word in s.
All the words in s are separated by a single space.
*********************************************/

func Output557() any {
	return reverseWords("Let's take LeetCode contest")
}

// * Solution 1
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := range words {
		var tmp string
		for j := len(words[i]) - 1; j >= 0; j-- {
			tmp += string(words[i][j])
		}
		words[i] = tmp
	}
	return strings.Join(words, " ")
}

// * Solution 2
// func reverseWords(s string) string {
// 	var tmp int
// 	length := len(s) - 1
// 	for i := 0; i <= length; i++ {
// 		if s[i] == ' ' || i == length {
// 			for j := i; j >= tmp; j-- {
// 				if s[j] != ' ' {
// 					s += string(s[j])
// 				}
// 				if i != length && j == tmp {
// 					s += " "
// 				}
// 				fmt.Println("Tmp is", tmp)
// 				fmt.Println("J is", j)
// 			}
// 			tmp = i
// 		}
// 	}
// 	fmt.Println("Res is", s[length+1:])
// 	return s[length+1:]
// }
