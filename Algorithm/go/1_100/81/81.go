package main

import "fmt"

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return nums[0] == target
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

type TestCase struct {
	nums   []int
	target int
}

func main() {
	testCases := []TestCase{
		{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 0},
	}
	for _, testCase := range testCases {
		result := search(testCase.nums, testCase.target)
		fmt.Println("nums:", testCase.nums, ", target:", testCase.target, ", result:", result)
	}
}
