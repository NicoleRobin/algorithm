package main

import (
	"fmt"
)

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		fmt.Println("left:", left, ", right:", right)
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			if nums[left] < nums[right] {
				break
			} else {
			}
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
		// {[]int{1, 3, 5}},
		// {[]int{2, 2, 2, 0, 1}},
		// {[]int{1, 1}},
		// {[]int{1, 3, 3}},
		{[]int{3, 3, 1, 3}},
		// {[]int{10, 1, 10, 10, 10}},
	}
	for _, testCase := range testCases {
		result := findMin(testCase.nums)
		fmt.Println("nums:", testCase.nums, ", result:", result)
	}
}
