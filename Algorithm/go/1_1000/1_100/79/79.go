package main

import (
	"fmt"
)

/*
思路: 回溯法
1、遍历二维数组，找到第一个字符相同的位置
2、从这个位置开始，向上下左右四个方向搜索，如果找到下一个字符，继续向四个方向搜索，直到找到最后一个字符
3、如果找到最后一个字符，返回true，否则返回false
4、如果四个方向都没有找到，返回false

注意：
1、需要一个visited数组，记录已经访问过的位置
*/

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	visited := make([][]bool, 0, len(board))
	for _, row := range board {
		visitedRow := make([]bool, len(row))
		visited = append(visited, visitedRow)
	}

	mLen := len(board)
	nLen := len(board[0])
	var backtrack func(i, j int, wordIndex int) bool
	backtrack = func(i, j int, wordIndex int) bool {
		// fmt.Println("i:", i, ", j:", j, ", wordIndex:", wordIndex)
		if wordIndex == len(word) {
			return true
		}

		if board[i][j] == word[wordIndex] {
			if wordIndex == len(word)-1 {
				return true
			}

			for _, direction := range directions {
				newi := i + direction[0]
				newj := j + direction[1]
				if newi < 0 || newi >= mLen {
					continue
				}
				if newj < 0 || newj >= nLen {
					continue
				}
				if visited[newi][newj] {
					continue
				}
				visited[newi][newj] = true
				if backtrack(newi, newj, wordIndex+1) {
					return true
				}
				visited[newi][newj] = false
			}
		}
		return false
	}

	for i := 0; i < mLen; i++ {
		for j := 0; j < nLen; j++ {
			visited[i][j] = true
			if backtrack(i, j, 0) {
				return true
			}
			visited[i][j] = false
		}
	}

	return false
}

type TestCase struct {
	board [][]byte
	word  string
}

func main() {
	testCases := []TestCase{
		/*
				{
					board: [][]byte{
						{'A', 'B', 'C', 'E'},
						{'S', 'F', 'C', 'S'},
						{'A', 'D', 'E', 'E'},
					},
					word: "ABCCED",
				},
				{
					board: [][]byte{
						{'A', 'B', 'C', 'E'},
						{'S', 'F', 'C', 'S'},
						{'A', 'D', 'E', 'E'},
					},
					word: "SEE",
				},
				{
					board: [][]byte{
						{'A', 'B', 'C', 'E'},
						{'S', 'F', 'C', 'S'},
						{'A', 'D', 'E', 'E'},
					},
					word: "ABCB",
				},
			{
				board: [][]byte{
					{'a'},
				},
				word: "a",
			},
		*/
		{
			board: [][]byte{
				{'a', 'a'},
			},
			word: "aaa",
		},
	}
	for _, testCase := range testCases {
		fmt.Println("word:", testCase.word, ", result:", exist(testCase.board, testCase.word))
	}
}
