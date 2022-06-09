package solutions

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}
	// 记录岛屿的数量
	numsOfLands := 0
	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				numsOfLands++
				// 传染
				infect(grid, i, j, row, col)
			}
		}
	}
	return numsOfLands
}

// infect 将[i][j]上下左右所有相连的1变为非1的值
func infect(grid [][]byte, i int, j int, row int, col int) {
	if grid == nil || i < 0 || i >= row || j < 0 || j >= col || grid[i][j] != '1' {
		return
	}
	// Infected 任何一个不等于1的值都可以
	const Infected = 2
	grid[i][j] = Infected
	infect(grid, i-1, j, row, col)
	infect(grid, i+1, j, row, col)
	infect(grid, i, j-1, row, col)
	infect(grid, i, j+1, row, col)
}
