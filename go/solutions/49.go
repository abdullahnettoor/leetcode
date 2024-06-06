package solutions

/* Question ********************************
? 49. Group Anagrams
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]

Constraints:
1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*****************************************/

func Output49() any {
	return groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}

// * Solution - Brute Force - Time O(n^2 × m) - Space O(n)
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	res := make([][]string, 0)
	addedWords := make(map[string]bool)
	for i := range strs {
		if addedWords[strs[i]] {
			continue
		}
		arr := []string{strs[i]}
		addedWords[strs[i]] = true
		for j := i + 1; j < len(strs); j++ {
			if isAnagram(strs[i], strs[j]) {
				arr = append(arr, strs[j])
				addedWords[strs[j]] = true
			}
		}
		res = append(res, arr)
	}
	return res
}
