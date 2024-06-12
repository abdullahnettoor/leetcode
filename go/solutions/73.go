package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

/* Question *******************************
? 73. Set Matrix Zeroes
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
You must do it in place.

Example 1:
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Example 2:
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

Constraints:
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1

Follow up:
A straightforward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
******************************************/

func Output73() any {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)

	return matrix
}

// * Solution -- Find Zero Indexes and Set Row and Col to Zeros -- Time O(m*n) - Space O(k) k=Number of Zeros in matrix
func setZeroes(matrix [][]int)  {
    zeroPosition := make(map[string]struct{})
    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] == 0 {
                key := fmt.Sprintf("%d-%d", i, j)
                zeroPosition[key] = struct{}{}
            }
        }
    }
    for k := range zeroPosition {
        rowCol := strings.Split(k, "-")
        row, _ := strconv.Atoi(rowCol[0])
        col, _ := strconv.Atoi(rowCol[1])

        for j := range matrix[row] {
          matrix[row][j] = 0
        }
        for i := range matrix {
            matrix[i][col] = 0
        }
    }
}
