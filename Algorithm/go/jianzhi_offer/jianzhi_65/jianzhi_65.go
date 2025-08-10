package main

import (
	"fmt"
)

func add(a int, b int) int {
	for b != 0 { // 当进位为 0 时跳出
		c := (a & b) << 1 // c = 进位
		a ^= b            // a = 非进位和
		b = c             // b = 进位
	}
	return a
}

func main() {
	fmt.Println("leetcode")
	a := 1
	b := 2
	sum := add(a, b)
	fmt.Printf("a:%d, b:%d, sum:%d\n", a, b, sum)
}
