package main

/*
思路：动态规划
类似于198，但是多了一个限制，即首尾不能同时被选中
所以需要分成两种情况来考虑
1、从0到n-2
2、从1到n-1
最后取两者的最大值
*/
func _rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	return max(_rob(nums[:n-1]), _rob(nums[1:]))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
