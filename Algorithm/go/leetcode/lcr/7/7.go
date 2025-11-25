package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
和LeetCode15是相同的；
需要考虑去重，
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)

	var ans [][]int
	for i := 0; i < len(nums); i++ {
		twoSumAns := twoSum(nums[i+1:], nums[i]*-1)
		if len(twoSumAns) > 0 {
			twoSumAns = append(twoSumAns, nums[i])
			ans = append(ans, twoSumAns)
		}
	}
	return ans
}

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		tmpSum := nums[left] + nums[right]
		if tmpSum == target {
			return []int{nums[left], nums[right]}
		}
		if tmpSum < target {
			left++
		}
		if tmpSum > target {
			right--
		}
	}
	return nil
}

func main() {
	testCases := []struct {
		nums     []int
		target   int
		expected [][]int
	}{
		{
			nums:     []int{-1, 0, 1, 2, -1, -4},
			target:   1,
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}
	for _, testCase := range testCases {
		result := threeSum(testCase.nums)
		if !reflect.DeepEqual(result, testCase.expected) {
			fmt.Printf("%d expected %d, got %d\n", testCase.nums, testCase.expected, result)
		} else {
			fmt.Printf("%d got %d\n", testCase.nums, testCase.expected)
		}
	}
}
