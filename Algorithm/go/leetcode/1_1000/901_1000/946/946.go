package main

// 栈、双指针
func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	var popIndex int
	for _, num := range pushed {
		stack = append(stack, num)
		for len(stack) > 0 && stack[len(stack)-1] == popped[popIndex] {
			stack = stack[:len(stack)-1]
			popIndex++
		}
	}

	if popIndex == len(popped) {
		return true
	}
	return false
}

func main() {}
