package main

import "fmt"

/*
思路：
1、利用单调栈，从底到顶存储的下标对应的temperature递减
*/
func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	// 单调栈，存储下标，从栈底到栈顶存储的下标对应的温度依次递减
	stack := []int{}
	for i, temperature := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 {
				if temperatures[stack[len(stack)-1]] < temperature {
					stackTopIndex := stack[len(stack)-1]
					result[stackTopIndex] = i - stackTopIndex
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			stack = append(stack, i)
		}
		{
			fmt.Printf("stack:")
			for _, index := range stack {
				fmt.Printf("%d ", temperatures[index])
			}
			fmt.Println()
		}
	}

	return result
}

func main() {
	nums := []int{73, 74, 75, 71, 69, 72, 76, 73}

	result := dailyTemperatures(nums)
	fmt.Println(result)
}
