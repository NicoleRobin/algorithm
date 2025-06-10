package main

import "fmt"

/*
思路：
1、将二维数组看成一维数组
2、使用两次二分查找
*/
func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	mLen := len(matrix)
	nLen := len(matrix[0])
	left, right := 0, mLen*nLen-1
	for left <= right {
		mid := (left + right) / 2
		xMid := mid / nLen
		yMid := mid % nLen
		if matrix[xMid][yMid] == target {
			return true
		} else if matrix[xMid][yMid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	mLen := len(matrix)
	left, right := 0, mLen-1

	// 先循环找到行
	var mIndex int
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid][0] <= target {
			mIndex = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// 在mIndex这一行进行二分查找，找目标值
	nLen := len(matrix[0])
	left, right = 0, nLen-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[mIndex][mid] == target {
			return true
		} else if matrix[mIndex][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	nums := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}

	// result := searchMatrix1(nums, 3)
	result := searchMatrix2(nums, 3)
	fmt.Println(result)
}
