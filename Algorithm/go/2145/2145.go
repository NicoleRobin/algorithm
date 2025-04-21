package main

import "math"

/*
思路：前缀和
但是实际上不用存储前缀和，只需要记录前缀和中的最大值和最小值即可；

然后设置一个start值，从lower到upper遍历，并分别计算start+maxPresum是否小于等于upper和start+minPresum是否大于等于lower
*/
func numberOfArrays(differences []int, lower int, upper int) int {
	var sum int
	maxSum, minSum := math.MinInt32, math.MaxInt32
	for _, diff := range differences {
		sum += diff
		if sum > maxSum {
			maxSum = sum
		}
		if sum < minSum {
			minSum = sum
		}
	}

	var result int
	for start := lower; start <= upper; start++ {
		if start+minSum >= lower && start+maxSum <= upper {
			result++
		}
	}

	return result
}

func main() {
}
