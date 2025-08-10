package main

/*
思考：
根据经验首先能想到的是用前缀和方法，但是具体做法，还需要再思考

思路：
1、将0和1分别转换为-1和1
*/
func findMaxLength(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	prefixSum := make(map[int]int)
	prefixSum[0] = -1
	maxLength := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if index, ok := prefixSum[sum]; ok {
			maxLength = max(maxLength, i-index)
		} else {
			prefixSum[sum] = i
		}
	}

	return maxLength
}

func main() {

}
