package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
			// 为什么right不是赋值为mid-1，因为mid可能是最小值
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right--
		}
	}

	return nums[left]
}

type TestCase struct {
	nums []int
}

func main() {
	fmt.Println("leetcode 154")
	testCases := []TestCase{
		{[]int{1, 3, 5}},
		{[]int{2, 2, 2, 0, 1}},
	}
	for _, testCase := range testCases {
		result := findMin(testCase.nums)
		fmt.Println("nums:", testCase.nums, ", result:", result)
	}
}
