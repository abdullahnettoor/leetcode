package solutions

/* Question **********************************
? 1089. Duplicate Zeros
Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.
Note that elements beyond the length of the original array are not written. Do the above modifications to the input array in place and do not return anything.

Example 1:
Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

Example 2:
Input: arr = [1,2,3]
Output: [1,2,3]
Explanation: After calling your function, the input array is modified to: [1,2,3]

Constraints:
1 <= arr.length <= 104
0 <= arr[i] <= 9
*********************************************/

func Output1089() any {
	arr := make([]int, 0)
	arr = append(arr, 1, 0, 2, 3, 0, 4, 5, 0)
	return duplicateZeros(arr)
}

func duplicateZeros(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			arr[i+1] = 0
			i++
		}
	}
	return arr
}

// * This is not valid answer, because the array can't be replaced 
// func duplicateZeros(arr []int) []int {
// 	var length = len(arr)
// 	for i := 0; i < length; i++ {
// 		fmt.Println("Array at", i, "is", arr)
// 		if arr[i] == 0 {
// 			tmp := arr[i+1:]
// 			fmt.Println("Tmp is", tmp)
// 			arr = append(arr[:i+2], tmp...)
// 			arr[i+1] = 0
// 			tmp = nil
// 			fmt.Println("Array ------", i, "is", arr)
// 			i++
// 		}
// 	}
// 	arr = arr[:length]
// 	return arr
// }`
