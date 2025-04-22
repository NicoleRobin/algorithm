package main

func countLargestGroup(n int) int {
	var maxCount int
	numSumCount := map[int]int{}
	for i := 1; i <= n; i++ {
		sum := numSum(i)
		numSumCount[sum]++
		if numSumCount[sum] > maxCount {
			maxCount = numSumCount[sum]
		}
	}

	var ans int
	for _, count := range numSumCount {
		if count == maxCount {
			ans++
		}
	}

	return ans
}

func numSum(num int) int {
	var sum int
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return sum
}

func main() {

}
