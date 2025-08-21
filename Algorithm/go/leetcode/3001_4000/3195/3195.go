package main

func minimumArea(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])

	var ans int
	letf, right, top, bottom := n, -1, m, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if j < letf {
					letf = j
				}
				if j > right {
					right = j
				}
				if i < top {
					top = i
				}
				if i > bottom {
					bottom = i
				}
			}
		}
	}

	return (max(0, right-letf+1) * max(0, bottom-top+1))
}

func main() {

}
