package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}

	partition := nums[l]
	i, j := l-1, r+1
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

func main() {
	fmt.Println("leetcode 215")
	nums := []int{3, 2, 1, 5, 6, 4}
	result := findKthLargest(nums, 2)
	fmt.Println(result)
}
