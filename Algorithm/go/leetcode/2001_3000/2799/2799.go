package main

import "fmt"

/*
思路：滑动窗口
1、使用map统计所有数字的个数
2、使用map统计当前窗口的数字个数
3、当当前窗口的数字个数等于所有数字的个数时，说明当前窗口是完整的子数组
*/
func countCompleteSubarrays(nums []int) int {
	allNumMap := map[int]int{}
	for _, num := range nums {
		allNumMap[num]++
	}
	allNumCount := len(allNumMap)

	left, right := 0, 0
	var result int
	numMap := map[int]int{}
	for left < len(nums) && right <= len(nums) {
		for len(numMap) < allNumCount && right < len(nums) {
			numMap[nums[right]]++
			right++
		}
		fmt.Println("numMap:", numMap, "left:", left, "right:", right)
		if len(numMap) == allNumCount {
			result += len(nums) - right + 1
			numMap[nums[left]]--
			if numMap[nums[left]] == 0 {
				delete(numMap, nums[left])
			}
			left++
		} else {
			break
		}
	}
	return result
}

func main() {
	testCases := []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{1, 3, 1, 2, 2},
			expect: 4,
		},
		{
			nums:   []int{5, 5, 5, 5},
			expect: 10,
		},
		{
			nums:   []int{459, 459, 962, 1579, 1435, 756, 1872, 1597},
			expect: 2,
		},
	}

	for i, testCase := range testCases {
		result := countCompleteSubarrays(testCase.nums)
		fmt.Println("i:", i, "result:", result, "expect:", testCase.expect)
	}
}
