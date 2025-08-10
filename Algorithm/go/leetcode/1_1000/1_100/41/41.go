package main

import "fmt"

/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

示例：
输入：nums = [1, 2, 0]
输出：3

输入：nums = [3, 4, -1, 1]
输出：2

输入：nums = [7, 8, 9, 11, 12]
输出：1
*/

/*
时间复杂度：O(n)
空间复杂度：O(n)
*/
func firstMissingPositive1(nums []int) int {
	numMap := map[int]interface{}{}
	max := 0
	for _, num := range nums {
		numMap[num] = struct{}{}
		if num > max {
			max = num
		}
	}

	for i := 1; i <= max; i++ {
		if _, ok := numMap[i]; !ok {
			return i
		}
	}
	return max + 1
}

/*
置换方法，将每个数字置换到正确的位置
时间复杂度：O(n)
空间复杂度：O(1)

3 4 -1 1
-1 4 3 1
-1 1 3 4
1 -1 3 4
*/
func firstMissingPositive2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			// 把num替换到正确的位置上nums[num-1]
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums)
}

func main() {
	fmt.Println("vim-go")
	testCases := [][]int{
		{1, 2, 0},
		{3, 4, -1, 1},
		{7, 8, 9, 11, 12},
	}
	for _, nums := range testCases {
		// result := firstMissingPositive1(nums)
		result := firstMissingPositive2(nums)
		fmt.Printf("nums:%v, result:%d\n", nums, result)
	}
}
