package main

func minOperations(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	if len(numSet) > 1 {
		return 1
	}
	return 0
}

func main() {}
