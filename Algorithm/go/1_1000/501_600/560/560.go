package main

import "fmt"

/*
思路：前缀和
前缀和数组preSums[i]表示nums[0:i-1]的和

参考：https://leetcode.cn/problems/range-sum-query-immutable/solutions/2693498/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/
*/
func subarraySum1(nums []int, k int) int {
	preSums := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		preSums[i] = preSums[i-1] + nums[i-1]
	}

	var ans int
	for i := 0; i < len(preSums); i++ {
		for j := i + 1; j < len(preSums); j++ {
			if preSums[j]-preSums[i] == k {
				ans++
			}
		}
	}
	return ans
}

/*
思路：前缀和+哈希表
preSumMap[preSum]表示preSum出现的次数
*/
func subarraySum(nums []int, k int) int {
	var ans int
	var preSum int
	preSumMap := map[int]int{
		0: 1, // preSum为0的情况，表示从头开始的子数组
	}
	for _, num := range nums {
		preSum += num
		ans += preSumMap[preSum-k] // 如果preSum-k在map中，说明存在一个子数组的和为k
		preSumMap[preSum]++
	}

	return ans
}

func main() {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2,
		},
		{
			nums:     []int{1},
			k:        0,
			expected: 0,
		},
	}
	for i, tc := range testCases {
		ans := subarraySum(tc.nums, tc.k)
		if ans != tc.expected {
			fmt.Println("Test case", i, "failed: got", ans, "expected", tc.expected)
		} else {
			fmt.Println("Test case", i, "passed")
		}
	}
}
