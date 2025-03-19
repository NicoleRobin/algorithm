package main

import (
	"fmt"
)

/*
思路：
两次二分查找，先找左边界，再找右边界
*/
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	leftIndex, rightIndex := -1, -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			leftIndex = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	left, right = 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			rightIndex = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{leftIndex, rightIndex}
}

func main() {
	fmt.Println("leetcode 34")
	nums := []int{5, 7, 7, 8, 8, 10}
	result := searchRange(nums, 8)
	fmt.Println(result)
}
