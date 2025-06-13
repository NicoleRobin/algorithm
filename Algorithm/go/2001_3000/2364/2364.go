package main

func countBadPairs(nums []int) int64 {
	var result int64
	diffMap := map[int]int{}
	for i, num := range nums {
		result += int64(i - diffMap[num-i])
		diffMap[num-i]++
	}
	return result
}

func main() {

}
