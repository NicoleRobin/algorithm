package main

import (
	"fmt"
	"strconv"
)

/*
思考：
1、低位在右边，需要从右边开始计算
2、用一个数字分别记录某一位的和，分别加上a/b以及carry
*/
func addBinary(a string, b string) string {
	ans := ""
	carry := 0

	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		ans = strconv.Itoa(sum%2) + ans
		carry = sum / 2
	}
	return ans
}

func main() {
	testCases := []struct {
		a, b   string
		expect string
	}{
		{
			a:      "11",
			b:      "1",
			expect: "100",
		},
	}

	for i, testCase := range testCases {
		ans := addBinary(testCase.a, testCase.b)
		if ans != testCase.expect {
			fmt.Printf("case %d fail, a:%s, b:%s, expect:%s, ans:%s\n", i, testCase.a, testCase.b, testCase.expect, ans)
		} else {
			fmt.Printf("case %d pass\n", i)
		}
	}
}
