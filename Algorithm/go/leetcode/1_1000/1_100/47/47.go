package main

import (
	"fmt"
	"reflect"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	// 将相同的数字排在一起
	sort.Ints(nums)
	var backtrack func(int, []int)
	var res [][]int
	visitedIndexSet := map[int]bool{}
	backtrack = func(index int, memo []int) {
		// fmt.Println("index:", index, ", memo:", memo, ", selectedIndexSet:", selectedIndexSet)
		if index == len(nums) {
			res = append(res, append([]int{}, memo...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if visitedIndexSet[i] {
				continue
			}
			if i > 0 && !visitedIndexSet[i-1] && nums[i] == nums[i-1] {
				// 跳过当前循环中选择过的重复的数字
				continue
			}
			visitedIndexSet[i] = true
			memo = append(memo, nums[i])
			backtrack(index+1, memo)
			visitedIndexSet[i] = false
			memo = memo[:len(memo)-1]
		}
	}
	backtrack(0, []int{})

	return res
}

func main() {
	testCases := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{1, 1, 2},
			expected: [][]int{},
		},
	}
	for i, tc := range testCases {
		res := permuteUnique(tc.nums)
		if reflect.DeepEqual(res, tc.expected) {
			fmt.Printf("Case %d fail, res:%v, expected:%v\n", i, res, tc.expected)
		} else {
			fmt.Printf("Case %d pass\n", i)
		}
	}
}
