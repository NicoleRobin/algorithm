package main

import (
	"fmt"
	"slices"
	"sort"
)

/*
思路：
排序，然后双指针
但是由于是一个区间，所以以往的方式不适用了；

lower−nums[j]≤nums[i]≤upper−nums[j]
*/
func countFairPairs(nums []int, lower, upper int) (ans int64) {
	slices.Sort(nums)
	for j, x := range nums {
		// upper-x+1中的+1是为了当遇到重复数字时，找到指定值的最后一个位置
		r := sort.SearchInts(nums[:j], upper-x+1)
		l := sort.SearchInts(nums[:j], lower-x)
		fmt.Println("j:", j, "x:", x, "l:", l, "r:", r)
		ans += int64(r - l)
	}
	return
}

/*
思路：
三指针
给数组 nums 排序后，随着枚举的 nums[j] 变大，upper−nums[j] 和 lower−nums[j] 都变小，于是我们可以用三指针来代替二分查找的搜索过程
*/
func countFairPairs2(nums []int, lower, upper int) int64 {
	slices.Sort(nums)
	left, right := len(nums), len(nums)
	var ans int64
	for j, x := range nums {
		for left > 0 && nums[left-1]+x >= lower {
			left--
		}
		for right > 0 && nums[right-1]+x > upper {
			right--
		}
		ans += int64(min(right, j) - min(left, j))
	}
	return ans
}

func main() {
	testCases := []struct {
		nums     []int
		lower    int
		upper    int
		expected int
	}{
		{
			nums:     []int{0, 1, 7, 4, 4, 5},
			lower:    3,
			upper:    6,
			expected: 6,
		},
	}

	for i, testCase := range testCases {
		result := countFairPairs(testCase.nums, testCase.lower, testCase.upper)
		println("Test Case", i+1, "Result:", result, "Expected:", testCase.expected)
	}
}
