package main

import "fmt"

func isValid(s string) bool {
	charMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for _, char := range s {
		if len(stack) > 0 {
			if stack[len(stack)-1] == charMap[byte(char)] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, byte(char))
			}
		} else {
			stack = append(stack, byte(char))
		}
	}

	return len(stack) == 0
}

func main() {
	s := "()"
	result := isValid(s)
	fmt.Println(result)
}
