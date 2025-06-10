package main

import "fmt"

/*
思考：
1、这里的难点是如何判断某一个点是否属于已有的岛屿

思路：
1、遍历二维数组，如果是1，就进行深度优先搜索，将所有相连的1都变成0
2、需要一个辅助grid，用来记录是否已经访问过
*/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	mLen := len(grid)
	nLen := len(grid[0])
	visited := make([][]bool, mLen)
	for i := 0; i < mLen; i++ {
		visited[i] = make([]bool, nLen)
	}

	var dfs func(i, j int)
	directions := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	dfs = func(i, j int) {
		grid[i][j] = '0'
		visited[i][j] = true
		for _, direction := range directions {
			newI := i + direction[0]
			newJ := j + direction[1]
			if newI < 0 || newI >= mLen || newJ < 0 || newJ >= nLen {
				continue
			}
			if grid[newI][newJ] == '1' && !visited[newI][newJ] {
				dfs(newI, newJ)
			}
		}
	}

	var result int
	for i := 0; i < mLen; i++ {
		for j := 0; j < nLen; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				result++
				dfs(i, j)
			}
		}
	}
	return result
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	result := numIslands(grid)
	fmt.Println(result)
}
