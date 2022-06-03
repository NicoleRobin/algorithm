package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	numMap := make(map[int]interface{})
	for _, num := range nums {
		numMap[num] = struct{}{}
	}

	result := []int{}
	for i := 1; i <= len(nums); i++ {
		if _, ok := numMap[i]; !ok {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	fmt.Println("vim-go")
}
