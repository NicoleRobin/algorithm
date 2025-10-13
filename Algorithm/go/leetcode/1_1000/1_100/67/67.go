package main

import "fmt"

/*
思考：
1、低位在右边，需要从右边开始计算
2、
*/
func addBinary(a string, b string) string {
}

func main() {
	testCases := []struct {
		a, b   string
		expect int
	}{
		{
			a:      "",
			b:      "",
			expect: -2,
		},
	}

	for i, testCase := range testCases {
		ans := addBinary(testCase.a, testCase.b)
		if ans != testCase.expect {
			fmt.Printf("case %d fail, a:%d, b:%d, expect:%d, ans:%d\n", i, testCase.a, testCase.b, testCase.expect, ans)
		} else {
			fmt.Printf("case %d pass\n", i)
		}
	}
}
