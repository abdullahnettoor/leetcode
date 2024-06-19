package solutions

/* Question **********************************
? 200. Number of Islands
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Example 2:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*********************************************/

func Output200() any {
	return numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	})
}

// * Solution -- BFS -- Time O(m*n) - Space O(m*n)
func numIslands(grid [][]byte) int {
	visited := make(map[[2]int]bool)
	islands := 0
	rows := len(grid)
	cols := len(grid[0])

	for r := range grid {
		for c := range grid[r] {
			pos := [2]int{r, c}
			if grid[r][c] == '1' && !visited[pos] {
				visited[pos] = true
				bfs := [][2]int{}
				bfs = append(bfs, pos)
				islands++
				for len(bfs) > 0 {
					curr := bfs[0]
					bfs = bfs[1:]
					for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
						r, c := d[0]+curr[0], d[1]+curr[1]
						pos := [2]int{r, c}
						if r >= 0 && r < rows && c >= 0 && c < cols &&
							grid[r][c] == '1' &&
							!visited[pos] {
							bfs = append(bfs, pos)
							visited[pos] = true
						}
					}

				}
			}
		}
	}
	return islands
}
