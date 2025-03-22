package main

import "fmt"

type Point struct {
	X int
	Y int
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func isValid(path []string, n int) bool {
	statuss := []Point{}
	for i, status := range path {
		for j, char := range status {
			if char == rune('Q') {
				statuss = append(statuss, Point{i, j})
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if statuss[i].Y == statuss[j].Y {
				return false
			}
			if abs(statuss[i].Y-statuss[j].Y) == abs(statuss[i].X-statuss[j].X) {
				return false
			}
		}
	}
	return true
}

func backtrack(row, n int, path []string, result *[][]string, colMap map[int]string) {
	if row == n {
		if isValid(path, n) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			*result = append(*result, tmp)
		}
		return
	}

	for col := 0; col < n; col++ {
		path = append(path, colMap[col])
		if isValid(path, row+1) {
			backtrack(row+1, n, path, result, colMap)
		}
		path = path[:len(path)-1]
	}
}

func solveNQueens(n int) [][]string {
	path := []string{}
	result := [][]string{}
	colMap := map[int]string{}
	for col := 0; col < n; col++ {
		status := ""
		for row := 0; row < n; row++ {
			if row == col {
				status += "Q"
			} else {
				status += "."
			}
		}
		colMap[col] = status
	}
	backtrack(0, n, path, &result, colMap)

	return result
}

func main() {
	result := solveNQueens(4)
	for _, res := range result {
		for _, status := range res {
			fmt.Println(status)
		}
		fmt.Println()
	}
}
