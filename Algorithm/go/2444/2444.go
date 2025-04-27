package main

import "math"

/*
思考：
1、直观的想法：双层遍历，枚举所有子数组，并在内层循环中遍历子数组并记录最大值、最小值，然后和指定的值比对；
2、但是，很明显上面的算法虽然是对的，但是效率很低，时间复杂度为O(n^3);
3、那么剩下的问题就是：如何快速的求一个子数组的最大值、最小值；
4、
*/
func countSubarrays(nums []int, minK int, maxK int) int64 {
	var ans int64
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			minNum, maxNum := math.MaxInt32, math.MinInt32
			for k := i; k <= j; k++ {
				if nums[k] < minNum {
					minNum = nums[k]
				}

				if nums[k] > maxNum {
					maxNum = nums[k]
				}
			}

			if minNum == minK && maxNum == maxK {
				ans++
			}
		}
	}
	return ans
}

func main() {

}
