package solutions

func SpiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	row, col := len(matrix), len(matrix[0])
	visited := make([][]bool, row, row)
	for i := range visited {
		visited[i] = make([]bool, col, col)
	}
	direct := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	rowIndex, colIndex, directIndex := 0, 0, 0
	for i := 0; i < row*col; i++ {
		visited[rowIndex][colIndex] = true
		res = append(res, matrix[rowIndex][colIndex])
		// 更新下一个需要遍历的元素
		nextRow, nextCol := rowIndex+direct[directIndex][0], colIndex+direct[directIndex][1]
		// 下一个要遍历的元素如果越界
		if nextRow < 0 || nextRow >= row || nextCol < 0 || nextCol >= col || visited[nextRow][nextCol] {
			directIndex = (directIndex + 1) % 4
		}
		rowIndex, colIndex = rowIndex+direct[directIndex][0], colIndex+direct[directIndex][1]
	}
	return res
}
