package solutions

/* Question **********************************
? 905. Sort Array By Parity
Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
Return any array that satisfies this condition.

Example 1:
Input: nums = [3,1,2,4]
Output: [2,4,3,1]
Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

Example 2:
Input: nums = [0]
Output: [0]

Constraints:
1 <= nums.length <= 5000
0 <= nums[i] <= 5000
*********************************************/

func Output905() any {
	return sortArrayByParity([]int{3, 1, 2, 4, 4, 6, 3, 1, 2})
}

// * Solution 1
func sortArrayByParity(nums []int) []int {
	even, odd := 0, len(nums)-1
	arr := make([]int, len(nums))
	for i := range nums {
		if nums[i]%2 == 0 {
			arr[even] = nums[i]
			even++
		} else {
			arr[odd] = nums[i]
			odd--
		}
	}
	return arr
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
