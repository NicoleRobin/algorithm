package main

import "fmt"

/*
509. 斐波那契数
时间复杂度：O(n)
*/
func fib(n int) int {
	if n <= 1 {
		return n
	}
	fn1, fn2 := 0, 1
	for i := 2; i <= n; i++ {
		tmp := fn1 + fn2
		fn1 = fn2
		fn2 = tmp
	}
	return fn2
}

func main() {
	fmt.Println(fib(4))
}
