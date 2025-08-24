package main

func partitionArray(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}
	arrayCount := len(nums) / k

	numCountMap := make(map[int]int)
	for _, num := range nums {
		numCountMap[num]++
	}
	for _, count := range numCountMap {
		if count > arrayCount {
			return false
		}
	}

	return true
}

func main() {}
