package main

/*
https://leetcode.cn/problems/zero-array-transformation-i/solutions/3674711/ling-shu-zu-bian-huan-i-by-leetcode-solu-7q94/
思路：差分数组
1. 创建一个差分数组 deltaArray，长度为 nums 的长度加 1。
2. 对于每个查询 [left, right]，在 deltaArray 的 left 位置加 1，在 right + 1 位置减 1。
*/
func isZeroArray(nums []int, queries [][]int) bool {
	deltaArray := make([]int, len(nums)+1)
	for _, query := range queries {
		left, right := query[0], query[1]
		deltaArray[left] += 1
		deltaArray[right+1] -= 1
	}

	operationCounts := make([]int, len(deltaArray))
	var currentOperationCount int
	for i, delta := range deltaArray {
		currentOperationCount += delta
		operationCounts[i] = currentOperationCount
	}
	for i := 0; i < len(nums); i++ {
		if operationCounts[i] < nums[i] {
			return false
		}
	}

	return true
}

func main() {

}
