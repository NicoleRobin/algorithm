package main

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	var ans []int
	for left < right {
		tmpSum := numbers[left] + numbers[right]
		if tmpSum == target {
			ans = append(ans, left+1, right+1)
			break
		} else if tmpSum < target {
			left++
		} else {
			right--
		}
	}
	return ans
}

func main() {

}
