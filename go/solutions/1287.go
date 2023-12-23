package solutions

/* Question **********************************
? 1287. Element Appearing More Than 25% In Sorted Array
Given an integer array sorted in non-decreasing order, there is exactly one integer in the array that occurs more than 25% of the time, return that integer.

Example 1:
Input: arr = [1,2,2,6,6,6,6,7,10]
Output: 6

Example 2:
Input: arr = [1,1]
Output: 1

Constraints:
1 <= arr.length <= 104
0 <= arr[i] <= 105
*********************************************/

func Output1287() any {
	return findSpecialInteger([]int{15, 15, 21, 21, 32, 32, 33, 33, 33, 34, 35})
}

// * Solution 1
func findSpecialInteger(arr []int) int {
	prev, count := arr[0], 1
	for i := range arr {
		if arr[i] == prev && i != 0 {
			count++
		} else {
			count = 1
		}
		if count > len(arr)/4 {
			break
		}
		prev = arr[i]
	}
	return prev
}
