package main

func findLHS(nums []int) int {
	numCountMap := map[int]int{}
	for _, num := range nums {
		numCountMap[num]++
	}

	var ans int
	for num, count := range numCountMap {
		if numCountMap[num-1] > 0 && count+numCountMap[num-1] > ans {
			ans = count + numCountMap[num-1]
		}
		if numCountMap[num+1] > 1 && count+numCountMap[num+1] > ans {
			ans = count + numCountMap[num+1]
		}
	}
	return ans
}

func main() {
}
