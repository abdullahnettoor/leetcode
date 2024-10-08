package solutions

/* Question **********************************
? 11. Container With Most Water
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.

Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
Input: height = [1,1]
Output: 1


Constraints:
n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*********************************************/

func Output11() any {
	return maxArea([]int{1, 3, 5, 4, 0, 1, 10})
}

// * Solution -- Two Pointers -- Time O(n) -- Space O(1)
func maxArea(height []int) int {
    var mostWater int
    start, end := 0, len(height)-1
    for start < end {
        w := end - start
        h := min(height[start], height[end])
        mostWater = max(mostWater, w * h)
        if height[start] < height[end] {
            start++
        } else {
            end--
        }
    }
    return mostWater
}

// * Solution -- Brute Force -- Time O(n^2) -- Space O(1)
// func maxArea(height []int) int {
//     var mostWater int
//     for i := 0; i < len(height)-1; i++ {
//         for j := len(height)-1; j > i; j-- {
//             h := min(height[i], height[j])
//             w := j - i
//             mostWater = max(mostWater, h * w)
//         }
//     }
//     return mostWater
// }

// func maxArea(height []int) int {
// 	var mostWater int
// 	start, end := 0, len(height)-1
// 	for start < end {
// 		for end > start {
// 			w := end - start
// 			h := min(height[start], height[end])
// 			mostWater = max(mostWater, w*h)
// 			end--
// 		}
// 		end = len(height) - 1
// 		start++
// 	}
// 	return mostWater
// }
