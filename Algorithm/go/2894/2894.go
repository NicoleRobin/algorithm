package main

import "fmt"

func differenceOfSum(n int, m int) int {
	sumOfAllNum := (1 + n) * n / 2
	// fmt.Println("sumOfAllNum:", sumOfAllNum)
	var num1, num2 int
	for i := 1; i <= n; i++ {
		if i*m > n {
			break
		}
		num2 += i * m
	}
	// fmt.Println("num2:", num2)
	num1 = sumOfAllNum - num2

	return num1 - num2
}

func main() {
	testCases := []struct {
		n        int
		m        int
		expected int
	}{
		{
			n:        10,
			m:        3,
			expected: 19,
		},
	}
	for i, testCase := range testCases {
		result := differenceOfSum(testCase.n, testCase.m)
		fmt.Println("i:", i, "n:", testCase.n, "m:", testCase.m, "result:", result, "expected:", testCase.expected)
	}
}
