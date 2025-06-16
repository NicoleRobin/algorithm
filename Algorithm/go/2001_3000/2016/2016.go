package main

import "fmt"

func maximumDifference(nums []int) int {
	ans := -1
	minValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minValue {
			minValue = nums[i]
		}
		if minValue < nums[i] && nums[i]-minValue > ans {
			ans = nums[i] - minValue
		}
	}
	return ans
}

func maximumDifference2(nums []int) int {
	var stack []int
	ans := -1
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && stack[len(stack)-1] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 && nums[i]-stack[0] > ans {
			ans = nums[i] - stack[0]
		}
		stack = append(stack, nums[i])
		fmt.Println("i:", i, "stack:", stack, "ans:", ans)
	}
	return ans
}

func maximumDifference1(nums []int) int {
	ans := -1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if nums[i]-nums[j] > ans {
					ans = nums[i] - nums[j]
				}
			}
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{7, 1, 5, 4},
			expected: 4,
		},
		{
			nums:     []int{9, 4, 3, 2},
			expected: -1,
		},
	}

	for i, tc := range testCases {
		ans := maximumDifference(tc.nums)
		if ans != tc.expected {
			println("Test case", i, "failed: expected", tc.expected, "but got", ans)
		} else {
			println("Test case", i, "passed")
		}
	}
}
