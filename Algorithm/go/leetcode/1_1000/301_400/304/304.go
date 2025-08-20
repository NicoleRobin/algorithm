package main

/*
思路：
1、二维前缀和
*/

type NumMatrix [][]int

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i, row := range matrix {
		sum[i+1] = make([]int, n+1)
		for j, x := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + x
		}
	}
	return sum
}

func (this NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this[row2+1][col2+1] - this[row2+1][col1] - this[row1][col2+1] + this[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {}
