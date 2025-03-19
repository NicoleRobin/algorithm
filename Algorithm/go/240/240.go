package main

import "fmt"

/*
思路：
每一行都是升序的，所以可以遍历所有行，在每一行中执行一次二分查找
时间复杂度：O(mlogn)
*/
func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	mLen := len(matrix)
	nLen := len(matrix[0])

	var result bool
	for mi := 0; mi < mLen; mi++ {
		left, right := 0, nLen-1
		for left <= right {
			mid := (left + right) / 2
			if matrix[mi][mid] == target {
				result = true
				break
			} else if matrix[mi][mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return result
}

/*
思路：
1、从右上角开始查找，如果当前值大于目标值，则向左移动一位，如果当前值小于目标值，则向下移动一位
2、直到找到目标值或越界
时间复杂度：O(m+n)

为什么选择右上角开始查找？
因为右上角的元素是这一行最大的，这一列最小的，所以可以根据这个元素的大小来判断目标值在这一行还是这一列
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	mLen := len(matrix)
	nLen := len(matrix[0])

	var result bool
	x, y := 0, nLen-1
	for x < mLen && y >= 0 {
		if matrix[x][y] == target {
			result = true
			break
		} else if matrix[x][y] > target {
			y--
		} else if matrix[x][y] < target {
			x++
		}
	}
	return result
}

func main() {
	matrix := [][]int{}

	result := searchMatrix(matrix, 1000)
	fmt.Println(result)
}
