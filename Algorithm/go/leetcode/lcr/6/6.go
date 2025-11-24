package main

func twoSum(nums []int, target int) []int {
	var ans []int
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == target {
			ans = append(ans, left, right)
			break
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			right--
		}
	}
	return ans
}

func main() {}
