package main

import (
	"fmt"
	"math"
	"sort"
)

/*
思路：贪心
1、将数组从大到小排序
2、以k为步长遍历

注意：
1、子数组 是数组中连续的 非空 元素序列。
所以是不能排序的
*/
func maxSubarraySum1(nums []int, k int) int64 {
	var ans int64
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums)
	for i := 0; i <= len(nums)-k; i = i + k {
		var tmp int64
		for j := i; j < i+k; j++ {
			tmp += int64(nums[j])
		}
		if ans != 0 && tmp <= 0 {
			break
		}
		ans += tmp
	}

	return ans
}

/*
思路：前缀和？
子数组和 =》前缀和
1、简单的二层循环会超时，怎么优化呢？
2、利用同余，计算并维护最小值
*/
func maxSubarraySum(nums []int, k int) int64 {
	ans := math.MinInt
	preSumNums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSumNums[i+1] = preSumNums[i] + nums[i]
	}

	minS := make([]int, k)
	for i := range minS {
		minS[i] = math.MaxInt / 2 // 防止下面减法溢出
	}

	for j, preSum := range preSumNums {
		i := j % k
		ans = max(ans, preSum-minS[i])
		minS[i] = min(minS[i], preSum)
	}

	return int64(ans)
}

func main() {
	testCases := []struct {
		nums   []int
		k      int
		expect int64
	}{
		{
			nums:   []int{1, 2},
			k:      1,
			expect: 3,
		},
		{
			nums:   []int{-5, 1, 2, -3, 4},
			k:      2,
			expect: 4,
		},
	}

	for i, testCase := range testCases {
		ans := maxSubarraySum(testCase.nums, testCase.k)
		if ans != testCase.expect {
			fmt.Printf("case %d fail, expect:%d, ans:%d\n", i, testCase.expect, ans)
		} else {
			fmt.Printf("case %d pass\n", i)
		}
	}
}
