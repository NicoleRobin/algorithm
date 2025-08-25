package main

import (
	"fmt"
	"reflect"
)

/*
官方题解
一共有 m+n−1 条对角线，相邻的对角线的遍历方向不同，当前遍历方向为从左下到右上，则紧挨着的下一条对角线遍历方向为从右上到左下；
设对角线从上到下的编号为 i∈[0,m+n−2]：
当 i 为偶数时，则第 i 条对角线的走向是从下往上遍历；
当 i 为奇数时，则第 i 条对角线的走向是从上往下遍历；
当第 i 条对角线从下往上遍历时，每次行索引减 1，列索引加 1，直到矩阵的边缘为止：
当 i<m 时，则此时对角线遍历的起点位置为 (i,0)；
当 i≥m 时，则此时对角线遍历的起点位置为 (m−1,i−m+1)；
当第 i 条对角线从上往下遍历时，每次行索引加 1，列索引减 1，直到矩阵的边缘为止：
当 i<n 时，则此时对角线遍历的起点位置为 (0,i)；
当 i≥n 时，则此时对角线遍历的起点位置为 (i−n+1,n−1)；
作者：力扣官方题解
链接：https://leetcode.cn/problems/diagonal-traverse/solutions/1597961/dui-jiao-xian-bian-li-by-leetcode-soluti-plz7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ans := make([]int, 0, m*n)
	for i := 0; i < m+n-1; i++ {
		if i%2 == 1 {
			x := max(i-n+1, 0)
			y := min(i, n-1)
			for x < m && y >= 0 {
				ans = append(ans, mat[x][y])
				x++
				y--
			}
		} else {
			x := min(i, m-1)
			y := max(i-m+1, 0)
			for x >= 0 && y < n {
				ans = append(ans, mat[x][y])
				x--
				y++
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func findDiagonalOrder1(mat [][]int) []int {
	if len(mat) == 0 {
		return []int{}
	}

	var ans []int
	m := len(mat)
	n := len(mat[0])
	for sum := 0; sum < m+n-1; sum++ {
		for i := 0; i < sum; i++ {
			ans = append(ans, mat[i][sum-i])
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		mat      [][]int
		expected []int
	}{
		{
			mat:      [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
	}
	for i, tc := range testCases {
		ans := findDiagonalOrder(tc.mat)
		if !reflect.DeepEqual(ans, tc.expected) {
			fmt.Printf("Case %d, fail, expected %d, got %d\n", i, tc.expected, ans)
		} else {
			fmt.Printf("Case %d, pass, expected %d, got %d\n", i, tc.expected, ans)
		}
	}
}
