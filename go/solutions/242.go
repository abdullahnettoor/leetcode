package solutions

import (
	"reflect"
	"sort"
	"strings"
)

/* Question ********************************
242. Valid Anagram
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Constraints:
1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.
*******************************************/

// * The following solution is not efficent. Will be update with new solution later.
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sChars := strings.Split(s, "")
	tChars := strings.Split(t, "")
	sort.Strings(sChars)
	sort.Strings(tChars)

	return reflect.DeepEqual(sChars, tChars)
}

// func contains(s []rune, e rune) bool {
// 	for _, a := range s {
// 		if a == e {
// 			return true
// 		}
// 	}
// 	return false
// }

// func IsAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	sChars := []rune(s)
// 	tChars := []rune(t)

// 	length := len(sChars)
// 	for i := 0; i < length; i++ {
// 		fmt.Println("len is", length)
// 		fmt.Println("i is", i)
// 		if ok := contains(sChars, tChars[i]); ok {
// 			fmt.Println("i", length)
// 			fmt.Println("sChar is", sChars)
// 			sChars = append(sChars[:i], sChars[(i+1):]...)
// 			length -= 1
// 			if length != 1 {
// 				i -= 1
// 			}
// 			fmt.Println("sChar is", sChars)
// 		}
// 	}
// 	fmt.Println("\nLength is", len(sChars))
// 	fmt.Println("sChar is", sChars)

// 	return len(sChars) == 1
// }
