package main

import (
	"fmt"
	"sort"
)

/*
分四种情况：
1、三个数都是0
2、一正一负一0
3、两正一负
4、两负一正

假如正数个数为x，负数个数为y
时间复杂度：O(x*y)
空间复杂度：O(n)
*/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	positiveNumMaps := map[int]interface{}{}
	negativeNumMaps := map[int]interface{}{}
	positiveNums := []int{}
	negativeNums := []int{}
	zeroNums := []int{}
	for _, num := range nums {
		if num == 0 {
			zeroNums = append(zeroNums, num)
		} else if num < 0 {
			negativeNumMaps[num] = struct{}{}
			negativeNums = append(negativeNums, num)
		} else {
			positiveNumMaps[num] = struct{}{}
			positiveNums = append(positiveNums, num)
		}
	}

	result := [][]int{}
	if len(zeroNums) >= 3 {
		result = append(result, []int{0, 0, 0})
	}

	if len(zeroNums) > 0 {
		for positiveNum, _ := range positiveNumMaps {
			negativeNum := positiveNum * -1
			if _, ok := negativeNumMaps[negativeNum]; ok {
				result = append(result, []int{0, positiveNum, negativeNum})
			}
		}
	}

	sort.Ints(negativeNums)
	sort.Ints(positiveNums)
	for num, _ := range negativeNumMaps {
		twoSumResults := twoSum(positiveNums, num)
		result = append(result, twoSumResults...)
	}

	for num, _ := range positiveNumMaps {
		twoSumResults := twoSum(negativeNums, num)
		result = append(result, twoSumResults...)
	}

	return result
}

func twoSum(nums []int, target int) [][]int {
	fmt.Printf("nums:%+v, target:%d\n", nums, target)
	target = target * -1
	result := [][]int{}
	left, right := 0, len(nums)-1
	for left < right {
		if left > 0 && nums[left] == nums[left-1] {
			left++
			continue
		}
		if right < len(nums)-1 && nums[right] == nums[right+1] {
			right--
			continue
		}
		sum := nums[left] + nums[right]
		fmt.Printf("left:%d, right:%d, sum:%d\n", left, right, sum)
		if sum == target {
			result = append(result, []int{nums[left], nums[right], target * -1})
			left++
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return result
}

func main() {
	fmt.Println("leetcode-15")
	nums := []int{-2, 0, 1, 1, 2}
	result := threeSum(nums)
	fmt.Printf("result:%+v\n", result)
}
