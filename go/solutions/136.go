package solutions

/* Question ********************************
7. Reverse Integer
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321

Example 3:
Input: x = 120
Output: 21

Constraints:
-231 <= x <= 231 - 1
*****************************************/

func Output136() any {
	return singleNumber([]int{3, 4, 3, 4, 5})
}

// * Solution -- Using Bitwise XOR Operator -- Time O(n) - Space O(1)
func singleNumber(nums []int) int {
    var res int
    for _, num := range nums {
        res = res ^ num
    }
    return res
}

// * Solution -- Hashmap -- Time & Space O(n)
// func singleNumber(nums []int) int {
// 	foundNTimes := make(map[int]int)
// 	for _, num := range nums {
// 		foundNTimes[num]++
// 	}
// 	var res int
// 	for k, v := range foundNTimes {
// 		if v == 1 {
// 			res = k
// 			break
// 		}
// 	}
// 	return res
// }

// * Solution -- Sorting -- Time O(n log n) - Space O(1)
// func singleNumber(nums []int) int {
//     if len(nums) == 1 {
//         return nums[0]
//     }
//     sort.Ints(nums)
//     ptr1, ptr2 := 0, 1 
//     fmt.Println(nums)
//     for ptr2 < len(nums) {
//         if nums[ptr1] != nums[ptr2]{
//             break
//         }
//         ptr1 += 2
//         ptr2 += 2
//     }
//     return nums[ptr1]
// }

// * Solution -- Brute Force -- Using Nested Loops -- Time O(n^2) - Space O(1)
// func singleNumber(nums []int) int {
// 	for i := 0; i < len(nums); i++ {
// 		count := 0
// 		for j := 0; j < len(nums); j++ {
// 			if nums[i] == nums[j] {
// 				count++
// 			}
// 		}
// 		if count == 1 {
// 			return nums[i]
// 		}
// 	}
// 	return -1
// }



