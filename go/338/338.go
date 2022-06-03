package main

import "fmt"

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = countBit(i)
	}

	return result
}

func countBit2(n int) int {
	count := 0
	for n > 0 {
		n &= (n - 1)
		count++
	}
	return count
}

func countBit1(n int) int {
	fmt.Println(n)
	count := 0
	for n > 0 {
		if n&1 > 0 {
			count++
		}
		n >>= 1
	}
	return count
}

func main() {
	fmt.Println("vim-go")
}
