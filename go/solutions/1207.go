package solutions

/* Question **********************************
? 1207. Unique Number of Occurrences
Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.

Example 1:
Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.

Example 2:
Input: arr = [1,2]
Output: false

Example 3:
Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true

Constraints:
1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000
*********************************************/

func Output1207() any {
	return uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})
}

// * Solution 1
func uniqueOccurrences(arr []int) bool {
	m1 := map[int]int{}
	for i := range arr {
		m1[arr[i]]++
	}
	m2 := map[int]bool{}
	for _, v := range m1 {
		if m2[v] {
			return false
		}
		m2[v] = true
	}
	return true
}
