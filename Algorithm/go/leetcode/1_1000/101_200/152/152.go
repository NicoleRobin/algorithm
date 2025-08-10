package main

import "fmt"

/*
动态规划
1、状态定义：
dpMax[i]：以nums[i]结尾的连续子数组最大乘积
dpMin[i]：以nums[i]结尾的连续子数组最小乘积
2、状态转移方程：
当nums[i]为正数时：
dpMax[i] = max(dpMax[i-1]*nums[i], nums[i])
dpMin[i] = min(dpMin[i-1]*nums[i], nums[i])
当nums[i]为负数时：
dpMax[i] = max(dpMin[i-1]*nums[i], nums[i])
dpMin[i] = min(dpMax[i-1]*nums[i], nums[i])
*/
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dpMin := make([]int, len(nums))
	dpMax := make([]int, len(nums))
	dpMin[0] = nums[0]
	dpMax[0] = nums[0]
	result := dpMax[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			if dpMax[i-1]*nums[i] > nums[i] {
				dpMax[i] = dpMax[i-1] * nums[i]
			} else {
				dpMax[i] = nums[i]
			}

			if dpMin[i-1]*nums[i] < nums[i] {
				dpMin[i] = dpMin[i-1] * nums[i]
			} else {
				dpMin[i] = nums[i]
			}
		} else if nums[i] < 0 {
			if dpMin[i-1]*nums[i] > nums[i] {
				dpMax[i] = dpMin[i-1] * nums[i]
			} else {
				dpMax[i] = nums[i]
			}

			if dpMax[i-1]*nums[i] < nums[i] {
				dpMin[i] = dpMax[i-1] * nums[i]
			} else {
				dpMin[i] = nums[i]
			}
		} else {
			dpMin[i] = 0
			dpMin[i] = 0
		}
		if dpMax[i] > result {
			result = dpMax[i]
		}
	}
	fmt.Println(dpMin)
	fmt.Println(dpMax)
	return result
}

func main() {
	fmt.Println("vim-go")
	nums := []int{-1, -2, -9, -6}
	result := maxProduct(nums)
	fmt.Println(result)
}
