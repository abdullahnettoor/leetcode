package solutions

/* Question **********************************
? 2085. Count Common Words With One Occurrence
Given two string arrays words1 and words2, return the number of strings that appear exactly once in each of the two arrays.

Example 1:
Input: words1 = ["leetcode","is","amazing","as","is"], words2 = ["amazing","leetcode","is"]
Output: 2
Explanation:
- "leetcode" appears exactly once in each of the two arrays. We count this string.
- "amazing" appears exactly once in each of the two arrays. We count this string.
- "is" appears in each of the two arrays, but there are 2 occurrences of it in words1. We do not count this string.
- "as" appears once in words1, but does not appear in words2. We do not count this string.
Thus, there are 2 strings that appear exactly once in each of the two arrays.

Example 2:
Input: words1 = ["b","bb","bbb"], words2 = ["a","aa","aaa"]
Output: 0
Explanation: There are no strings that appear in each of the two arrays.

Example 3:
Input: words1 = ["a","ab"], words2 = ["a","a","a","ab"]
Output: 1
Explanation: The only string that appears exactly once in each of the two arrays is "ab".

Constraints:
1 <= words1.length, words2.length <= 1000
1 <= words1[i].length, words2[j].length <= 30
words1[i] and words2[j] consists only of lowercase English letters.
*********************************************/

func Output2085() any {
	return countWords([]string{"leetcode", "is", "amazing", "as", "is"}, []string{"amazing", "leetcode", "is"})
}

// * Solution 1 -- Using individual loops
func countWords(words1 []string, words2 []string) int {
	m := make(map[string]int, len(words1))
	ctr := 0
	for i := range words1 {
		m[words1[i]]++
	}
	for i := range words2 {
		if m[words2[i]] > 1 {
			m[words2[i]] = 0
		}
		m[words2[i]]--
	}
	for _, v := range m {
		if v == 0 {
			ctr++
		}
	}
	return ctr
}

// * Solution 2 -- Using nested loop
// func countWords(words1 []string, words2 []string) int {
// 	m := map[string]int{}
// 	ctr := 0
// 	for i := range words1 {
// 		for j := range words2 {
// 			if words1[i] == words2[j] {
// 				m[words1[i]]++
// 			}
// 		}
// 	}
// 	for _, v := range m {
// 		if v == 1 {
// 			ctr++
// 		}
// 	}
// 	return ctr
// }
