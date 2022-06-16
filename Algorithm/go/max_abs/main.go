package main

import (
	"fmt"
	"math"
)

func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	leftMax := make([]int, len(nums))
	rightMax := make([]int, len(nums))
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		leftMax[i] = max
	}

	max = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i+1] > max {
			max = nums[i+1]
		}
		rightMax[i] = max
	}
	fmt.Printf("leftMax:%v\n", leftMax)
	fmt.Printf("rightMax:%v\n", rightMax)

	result := math.MinInt32
	for i := 0; i < len(nums)-1; i++ {
		if abs(leftMax[i]-rightMax[i]) > result {
			result = abs(leftMax[i] - rightMax[i])
		}
	}
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}
func main() {
	fmt.Println("vim-go")
	nums := []int{2, 7, 3, 1, 1}
	result := max(nums)
	fmt.Println(result)
}
