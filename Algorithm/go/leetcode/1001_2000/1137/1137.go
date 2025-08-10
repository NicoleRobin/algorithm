package main

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}

	t1, t2, t3 := 0, 1, 1
	for i := 3; i <= n; i++ {
		t1, t2, t3 = t2, t3, t1+t2+t3
	}
	return t3
}
func main() {
	n := 4
	result := tribonacci(n)
	println(result) // Output: 4
}
