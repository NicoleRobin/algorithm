package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	var left int
	var sum int
	res := len(nums) - 1
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < res {
				res = right - left+1
			}
			sum -= nums[left]
			left++
		}
		fmt.Println("left:", left, "right:", right, "sum:", sum)
	}
	return res
}

func main() {
	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{2, 3, 1, 2, 4, 3},
			target:   7,
			expected: 2,
		},
	}
	for i, tc := range testCases {
		res := minSubArrayLen(tc.target, tc.nums)
		if res != tc.expected {
			fmt.Printf("#%d failed, res:%d, expected:%d\n", i, res, tc.expected)
		} else {
			fmt.Printf("#%d pass", i)
		}
	}
}
