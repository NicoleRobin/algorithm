package main

import "fmt"

func main() {
	testCases := []struct {
		a, b, expect int
	}{
		{
			a:      7,
			b:      -3,
			expect: -2,
		},
	}

	for i, testCase := range testCases {
		ans := divide(testCase.a, testCase.b)
		if ans != testCase.expect {
			fmt.Printf("case %d fail, a:%d, b:%d, expect:%d, ans:%d\n", i, testCase.a, testCase.b, testCase.expect, ans)
		} else {
			fmt.Printf("case %d pass\n", i)
		}
	}
}
