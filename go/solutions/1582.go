package solutions

/*  Question **********************************
? 1582. Special Positions in a Binary Matrix
Given an m x n binary matrix mat, return the number of special positions in mat.

A position (i, j) is called special if mat[i][j] == 1 and all other elements in row i and column j are 0 (rows and columns are 0-indexed).

Example 1:
Input: mat = [[1,0,0],[0,0,1],[1,0,0]]
Output: 1
Explanation: (1, 2) is a special position because mat[1][2] == 1 and all other elements in row 1 and column 2 are 0.

Example 2:
Input: mat = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3
Explanation: (0, 0), (1, 1) and (2, 2) are special positions.

Constraints:
m == mat.length
n == mat[i].length
1 <= m, n <= 100
mat[i][j] is either 0 or 1.
*********************************************/

func Output1582() any {
	return numSpecial([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 1}})
}

func numSpecial(mat [][]int) int {
	var count int
	for i := 0; i < len(mat); i++ {
		idx := checkRow(mat[i])
		if idx == -1 {
			continue
		}
		if checkColumn(mat, idx, i) {
			count++
		}
	}
	return count
}

func checkRow(arr []int) int {
	var tmp, index = 0, -1
	for i := range arr {
		if arr[i] == 0 {
			continue
		}
		tmp++
		if tmp > 1 {
			return -1
		}
		index = i
	}
	return index
}

func checkColumn(arr [][]int, idx, i int) bool {
	var tmp int
	for i := range arr {
		if arr[i][idx] == 0 {
			continue
		}
		tmp++
		if tmp > 1 {
			return false
		}
	}
	return true
}
