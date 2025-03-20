package main

/*
思路：
1、有点像是二叉树的层序遍历，只不过这里是四个方向
*/
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	mLen := len(grid)
	nLen := len(grid[0])

	var stack [][2]int
	var freshCount int
	for i := 0; i < mLen; i++ {
		for j := 0; j < nLen; j++ {
			if grid[i][j] == 2 {
				stack = append(stack, [2]int{i, j})
			}
			if grid[i][j] == 1 {
				freshCount++
			}
		}
	}
	if freshCount == 0 {
		return 0
	}

	directions := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	var loopCount int
	for len(stack) > 0 {
		loopCount++
		var nextStack [][2]int
		for _, position := range stack {
			for _, direction := range directions {
				newX := position[0] + direction[0]
				newY := position[1] + direction[1]
				if newX < 0 || newX >= mLen || newY < 0 || newY >= nLen {
					continue
				}
				if grid[newX][newY] == 1 {
					grid[newX][newY] = 2
					freshCount--
					nextStack = append(nextStack, [2]int{newX, newY})
				}
			}
		}
		stack = nextStack
	}

	if freshCount > 0 {
		return -1
	}
	return loopCount - 1
}

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	result := orangesRotting(grid)
	println(result)
}
